//go:build all || unit
// +build all unit

package hedera

/*-
 *
 * Hedera Go SDK
 *
 * Copyright (C) 2020 - 2022 Hedera Hashgraph, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/hashgraph/hedera-protobufs-go/services"
	protobuf "google.golang.org/protobuf/proto"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestUnitContractUpdateTransactionValidate(t *testing.T) {
	client := ClientForTestnet()
	client.SetAutoValidateChecksums(true)
	contractID, err := ContractIDFromString("0.0.123-esxsf")
	require.NoError(t, err)
	accountID, err := AccountIDFromString("0.0.123-esxsf")
	require.NoError(t, err)
	fileID, err := FileIDFromString("0.0.123-esxsf")
	require.NoError(t, err)

	contractInfoQuery := NewContractUpdateTransaction().
		SetContractID(contractID).
		SetProxyAccountID(accountID).
		SetBytecodeFileID(fileID)

	err = contractInfoQuery._ValidateNetworkOnIDs(client)
	require.NoError(t, err)
}

func TestUnitContractUpdateTransactionValidateWrong(t *testing.T) {
	client := ClientForTestnet()
	client.SetAutoValidateChecksums(true)
	contractID, err := ContractIDFromString("0.0.123-rmkykd")
	require.NoError(t, err)
	accountID, err := AccountIDFromString("0.0.123-rmkykd")
	require.NoError(t, err)
	fileID, err := FileIDFromString("0.0.123-rmkykd")
	require.NoError(t, err)

	contractInfoQuery := NewContractUpdateTransaction().
		SetContractID(contractID).
		SetProxyAccountID(accountID).
		SetBytecodeFileID(fileID)

	err = contractInfoQuery._ValidateNetworkOnIDs(client)
	assert.Error(t, err)
	if err != nil {
		assert.Equal(t, "network mismatch or wrong checksum given, given checksum: rmkykd, correct checksum esxsf, network: testnet", err.Error())
	}
}

func TestUnitMockContractUpdateTransaction(t *testing.T) {
	call := func(request *services.Transaction) *services.TransactionResponse {
		require.NotEmpty(t, request.SignedTransactionBytes)
		signedTransaction := services.SignedTransaction{}
		_ = protobuf.Unmarshal(request.SignedTransactionBytes, &signedTransaction)

		require.NotEmpty(t, signedTransaction.BodyBytes)
		transactionBody := services.TransactionBody{}
		_ = protobuf.Unmarshal(signedTransaction.BodyBytes, &transactionBody)

		require.NotNil(t, transactionBody.TransactionID)
		transactionId := transactionBody.TransactionID.String()
		require.NotEqual(t, "", transactionId)

		sigMap := signedTransaction.GetSigMap()
		require.NotNil(t, sigMap)

		for _, sigPair := range sigMap.SigPair {
			verified := false

			switch k := sigPair.Signature.(type) {
			case *services.SignaturePair_Ed25519:
				pbTemp, _ := PublicKeyFromBytesEd25519(sigPair.PubKeyPrefix)
				verified = pbTemp.Verify(signedTransaction.BodyBytes, k.Ed25519)
			case *services.SignaturePair_ECDSASecp256K1:
				pbTemp, _ := PublicKeyFromBytesECDSA(sigPair.PubKeyPrefix)
				verified = pbTemp.Verify(signedTransaction.BodyBytes, k.ECDSASecp256K1)
			}
			require.True(t, verified)
		}

		if bod, ok := transactionBody.Data.(*services.TransactionBody_ContractUpdateInstance); ok {
			require.Equal(t, bod.ContractUpdateInstance.ContractID.GetContractNum(), int64(3))
			if mem, ok2 := bod.ContractUpdateInstance.MemoField.(*services.ContractUpdateTransactionBody_MemoWrapper); ok2 {
				require.Equal(t, mem.MemoWrapper.GetValue(), "yes")
			}
			require.Equal(t, hex.EncodeToString(bod.ContractUpdateInstance.GetAdminKey().GetEd25519()), "1480272863d39c42f902bc11601a968eaf30ad662694e3044c86d5df46fabfd2")
		}

		return &services.TransactionResponse{
			NodeTransactionPrecheckCode: services.ResponseCodeEnum_OK,
		}
	}
	responses := [][]interface{}{{
		call,
	}}

	client, server := NewMockClientAndServer(responses)
	defer server.Close()
	//302a300506032b65700321001480272863d39c42f902bc11601a968eaf30ad662694e3044c86d5df46fabfd2
	newKey, err := PrivateKeyFromStringEd25519("302e020100300506032b657004220420278184257eb568d0e5fcfc1df99828b039b4776da05855dc5af105996e6200d1")
	require.NoError(t, err)

	tran := TransactionIDGenerate(AccountID{Account: 3})

	_, err = NewContractUpdateTransaction().
		SetNodeAccountIDs([]AccountID{{Account: 3}}).
		SetTransactionID(tran).
		SetAdminKey(newKey.PublicKey()).
		SetContractMemo("yes").
		SetContractID(ContractID{Contract: 3}).
		Execute(client)
	require.NoError(t, err)
}

func TestUnitContractUpdateTransactionGet(t *testing.T) {
	contractID := ContractID{Contract: 7}

	nodeAccountID := []AccountID{{Account: 10}, {Account: 11}, {Account: 12}}
	transactionID := TransactionIDGenerate(AccountID{Account: 324})

	newKey, err := PrivateKeyGenerateEd25519()

	transaction, err := NewContractUpdateTransaction().
		SetTransactionID(transactionID).
		SetNodeAccountIDs(nodeAccountID).
		SetContractID(contractID).
		SetAdminKey(newKey.PublicKey()).
		SetContractMemo("yes").
		SetTransactionMemo("").
		SetTransactionValidDuration(60 * time.Second).
		SetRegenerateTransactionID(false).
		Freeze()
	require.NoError(t, err)

	transaction.GetTransactionID()
	transaction.GetNodeAccountIDs()

	_, err = transaction.GetTransactionHash()
	require.NoError(t, err)

	transaction.GetContractID()
	transaction.GetMaxTransactionFee()
	transaction.GetTransactionMemo()
	transaction.GetRegenerateTransactionID()
	_, err = transaction.GetSignatures()
	require.NoError(t, err)
	transaction.GetRegenerateTransactionID()
	transaction.GetMaxTransactionFee()
	transaction.GetAdminKey()
	transaction.GetRegenerateTransactionID()
	transaction.GetContractMemo()
}

func TestUnitContractUpdateTransactionSetNothing(t *testing.T) {
	nodeAccountID := []AccountID{{Account: 10}, {Account: 11}, {Account: 12}}
	transactionID := TransactionIDGenerate(AccountID{Account: 324})

	transaction, err := NewContractUpdateTransaction().
		SetTransactionID(transactionID).
		SetNodeAccountIDs(nodeAccountID).
		Freeze()
	require.NoError(t, err)

	transaction.GetTransactionID()
	transaction.GetNodeAccountIDs()

	_, err = transaction.GetTransactionHash()
	require.NoError(t, err)

	transaction.GetContractID()
	transaction.GetMaxTransactionFee()
	transaction.GetTransactionMemo()
	transaction.GetRegenerateTransactionID()
	_, err = transaction.GetSignatures()
	require.NoError(t, err)
	transaction.GetRegenerateTransactionID()
	transaction.GetMaxTransactionFee()
	transaction.GetAdminKey()
	transaction.GetRegenerateTransactionID()
	transaction.GetContractMemo()
}
