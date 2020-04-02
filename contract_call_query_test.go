package hedera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContractCallQuery(t *testing.T) {
	mockTransaction, err := newMockTransaction()
	assert.NoError(t, err)

	parameters := NewContractFunctionParams().
		AddBytes([]byte{24, 43, 11})

	query := NewContractCallQuery().
		SetGas(100).
		SetMaxResultSize(100).
		SetFunction("someFunction", parameters).
		SetContractID(ContractID{Contract: 3}).
		SetQueryPaymentTransaction(mockTransaction)

	assert.Equal(t, `contractCallLocal:<header:<payment:<bodyBytes:"\n\016\n\010\010\334\311\007\020\333\237\t\022\002\030\003\022\002\030\003\030\200\302\327/\"\002\010xr\024\n\022\n\007\n\002\030\002\020\307\001\n\007\n\002\030\003\020\310\001" sigMap:<sigPair:<pubKeyPrefix:"\344\361\300\353L}\315\303\347\353\021p\263\010\212=\022\242\227\364\243\353\342\362\205\003\375g5F\355\216" ed25519:"\022&5\226\373\264\034]P\273%\354P\233k\315\231\013\337\274\254)\246+\322<\227+\273\214\212f\313\332i\027T4{\367\363UYn\n\217\253ep\004\366\203\017\272FUP\243\321/\035\235\032\013" > > > > contractID:<contractNum:3 > gas:100 functionParameters:"7Q\372\261\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000 \000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\003\030+\013\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000" maxResultSize:100 > `, query.QueryBuilder.pb.String())
}
