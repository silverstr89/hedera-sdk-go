package hedera

import (
	"strings"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

type ErrorTransactionValidation struct {
	Messages []string
	Err      error
}

func (e *ErrorTransactionValidation) Error() string {
	return "The following requirements were not met: \n" + strings.Join(e.Messages, "\n")
}

type TransactionBuilderInterface interface {
	Validate() error
	Build() (*Transaction, error)
	Execute() (*TransactionID, error)
	ExecuteForReceipt() (*TransactionReceipt, error)
}

type TransactionBuilder struct {
	client            *Client
	kind              TransactionKind
	MaxTransactionFee uint64
	body              proto.TransactionBody
}

func (tb TransactionBuilder) build() (*Transaction, error) {
	if tb.client != nil {
		if tb.body.TransactionFee == 0 {
			tb.body.TransactionFee = tb.client.MaxTransactionFee()
		}

		if tb.body.TransactionValidDuration == nil {
			tb.body.TransactionValidDuration = &proto.Duration{Seconds: maxValidDuration}
		}

		if tb.body.NodeAccountID == nil {
			// let the client pick an actual node
			tb.body.NodeAccountID = tb.client.nodeID.proto()
		}
	}

	if tb.MaxTransactionFee == 0 {
		if tb.client != nil {
			tb.body.TransactionFee = tb.MaxTransactionFee
		}
	}

	protoBody := proto.Transaction_Body{
		Body: &tb.body,
	}

	tx := Transaction{
		Kind:   tb.kind,
		client: tb.client,
		inner: proto.Transaction{
			BodyData: &protoBody,
		},
	}

	return &tx, nil
}
