package hedera

import (
	"context"
	"math"
	"time"

	"github.com/hashgraph/hedera-sdk-go/proto"
	"google.golang.org/grpc"
)

const maxAttempts = 10;

type method struct {
	query func(
		context.Context,
		*proto.Query,
		...grpc.CallOption,
	) (*proto.Response, error)
	transaction func(
		context.Context,
		*proto.Transaction,
		...grpc.CallOption,
	) (*proto.TransactionResponse, error)
}

type response struct {
	query       *proto.Response
	transaction *proto.TransactionResponse
}

type intermediateResponse struct {
	query       *proto.Response
	transaction TransactionResponse
}

type protoRequest struct {
	query       *proto.Query
	transaction *proto.Transaction
}

type QueryHeader struct {
	header *proto.QueryHeader
}

type protoResponseHeader struct {
	responseHeader proto.ResponseHeader
}

type request struct {
	query       *Query
	transaction *Transaction
}

func execute(
	client *Client,
	request request,
	shouldRetry func(Status, response) bool,
	makeRequest func(request) protoRequest,
	advanceRequest func(request),
	getNodeId func(request, *Client) NodeID,
	getMethod func(request, *channel) method,
	mapResponseStatus func(request, response) Status,
	mapResponse func(request, response, AccountID, protoRequest) (intermediateResponse, error),
) (intermediateResponse, error) {
	var attempt int64
	for attempt = 0; ; /* loop forever */ attempt++ {
		protoRequest := makeRequest(request)
		node := getNodeId(request, client)

		channel, err := client.getChannel(node.AccountID)
		if err != nil {
			return intermediateResponse{}, nil
		}

		method := getMethod(request, channel)

		advanceRequest(request)

		resp := response{}

		if !node.isHealthy(){
			node.wait()
		}

		if method.query != nil {
			r, err := method.query(context.TODO(), protoRequest.query)
			if err != nil {
				node.increaseDelay()
				continue
			}

			resp.query = r
		} else {
			r, err := method.transaction(context.TODO(), protoRequest.transaction)
			if err != nil {
				node.increaseDelay()
				continue
			}

			resp.transaction = r
		}

		node.decreaseDelay()

		status := mapResponseStatus(request, resp)

		if shouldRetry(status, resp) && attempt <= maxAttempts {
			delayForAttempt(attempt)
			continue
		}

		if status != StatusOk {
			return intermediateResponse{}, newErrHederaPreCheckStatus(TransactionID{}, status)
		}

		return mapResponse(request, resp, node.AccountID, protoRequest)
	}
}

func delayForAttempt(attempt int64) {
	// 0.1s, 0.2s, 0.4s, 0.8s, ...
	ms := int64(math.Floor(50 * math.Pow(2, float64(attempt))))
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
