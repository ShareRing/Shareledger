package pos

import (
	"fmt"
	"testing"
	// "time"

	// sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/sharering/shareledger/codec"
	"github.com/sharering/shareledger/types"
)

func TestValidatorEqual(t *testing.T) {
	val1 := NewValidator(addr1, pk1, Description{})
	val2 := NewValidator(addr1, pk1, Description{})

	ok := val1.Equal(val2)
	require.True(t, ok)

	val2 = NewValidator(addr2, pk2, Description{})

	ok = val1.Equal(val2)
	require.False(t, ok)
}

func TestABCIValidator(t *testing.T) {
	validator := NewValidator(addr1, pk1, Description{})

	abciVal := validator.ABCIValidator()
	require.Equal(t, tmtypes.TM2PB.PubKey(validator.GetABCIPubKey()), abciVal.PubKey)
	require.Equal(t, validator.BondedTokens().RoundInt64(), abciVal.Power)
}

func TestABCIValidatorZero(t *testing.T) {
	validator := NewValidator(addr1, pk1, Description{})

	abciVal := validator.ABCIValidatorZero()
	require.Equal(t, tmtypes.TM2PB.PubKey(validator.GetABCIPubKey()), abciVal.PubKey)
	require.Equal(t, int64(0), abciVal.Power)
}

func TestRemoveTokens(t *testing.T) {

	validator := Validator{
		Owner:           addr1,
		PubKey:          pk1,
		Status:          types.Bonded,
		Tokens:          types.NewDec(100),
		DelegatorShares: types.NewDec(100),
	}

	pool := InitialPool()
	pool.LooseTokens = types.NewDec(10)
	pool.BondedTokens = validator.BondedTokens()

	validator, pool = validator.UpdateStatus(pool, types.Bonded)
	require.Equal(t, types.Bonded, validator.Status)

	// remove tokens and test check everything
	validator, pool = validator.RemoveTokens(pool, types.NewDec(10))
	require.Equal(t, int64(90), validator.Tokens.RoundInt64())
	require.Equal(t, int64(90), pool.BondedTokens.RoundInt64())
	require.Equal(t, int64(20), pool.LooseTokens.RoundInt64())

	// update validator to unbonded and remove some more tokens
	validator, pool = validator.UpdateStatus(pool, types.Unbonded)
	require.Equal(t, types.Unbonded, validator.Status)
	require.Equal(t, int64(0), pool.BondedTokens.RoundInt64())
	require.Equal(t, int64(110), pool.LooseTokens.RoundInt64())

	validator, pool = validator.RemoveTokens(pool, types.NewDec(10))
	require.Equal(t, int64(80), validator.Tokens.RoundInt64())
	require.Equal(t, int64(0), pool.BondedTokens.RoundInt64())
	require.Equal(t, int64(110), pool.LooseTokens.RoundInt64())
}

func TestAddTokensValidatorBonded(t *testing.T) {
	pool := InitialPool()
	pool.LooseTokens = types.NewDec(10)
	validator := NewValidator(addr1, pk1, Description{})
	validator, pool = validator.UpdateStatus(pool, types.Bonded)
	validator, pool, delShares := validator.AddTokensFromDel(pool, types.NewInt(10))

	require.Equal(t, types.OneDec(), validator.DelegatorShareExRate())

	assert.True(types.DecEq(t, types.NewDec(10), delShares))
	assert.True(types.DecEq(t, types.NewDec(10), validator.BondedTokens()))
}

func TestAddTokensValidatorUnbonding(t *testing.T) {
	pool := InitialPool()
	pool.LooseTokens = types.NewDec(10)
	validator := NewValidator(addr1, pk1, Description{})
	validator, pool = validator.UpdateStatus(pool, types.Unbonding)
	validator, pool, delShares := validator.AddTokensFromDel(pool, types.NewInt(10))

	require.Equal(t, types.OneDec(), validator.DelegatorShareExRate())

	assert.True(types.DecEq(t, types.NewDec(10), delShares))
	assert.Equal(t, types.Unbonding, validator.Status)
	assert.True(types.DecEq(t, types.NewDec(10), validator.Tokens))
}

func TestAddTokensValidatorUnbonded(t *testing.T) {
	pool := InitialPool()
	pool.LooseTokens = types.NewDec(10)
	validator := NewValidator(addr1, pk1, Description{})
	validator, pool = validator.UpdateStatus(pool, types.Unbonded)
	validator, pool, delShares := validator.AddTokensFromDel(pool, types.NewInt(10))

	require.Equal(t, types.OneDec(), validator.DelegatorShareExRate())

	assert.True(types.DecEq(t, types.NewDec(10), delShares))
	assert.Equal(t, types.Unbonded, validator.Status)
	assert.True(types.DecEq(t, types.NewDec(10), validator.Tokens))
}

// TODO refactor to make simpler like the AddToken tests above
func TestRemoveDelShares(t *testing.T) {
	valA := Validator{
		Owner:           addr1,
		PubKey:          pk1,
		Status:          types.Bonded,
		Tokens:          types.NewDec(100),
		DelegatorShares: types.NewDec(100),
	}
	poolA := InitialPool()
	poolA.LooseTokens = types.NewDec(10)
	poolA.BondedTokens = valA.BondedTokens()
	require.Equal(t, valA.DelegatorShareExRate(), types.OneDec())

	// Remove delegator shares
	valB, poolB, coinsB := valA.RemoveDelShares(poolA, types.NewDec(10))
	assert.Equal(t, int64(10), coinsB.RoundInt64())
	assert.Equal(t, int64(90), valB.DelegatorShares.RoundInt64())
	assert.Equal(t, int64(90), valB.BondedTokens().RoundInt64())
	assert.Equal(t, int64(90), poolB.BondedTokens.RoundInt64())
	assert.Equal(t, int64(20), poolB.LooseTokens.RoundInt64())

	// conservation of tokens
	require.True(types.DecEq(t,
		poolB.LooseTokens.Add(poolB.BondedTokens),
		poolA.LooseTokens.Add(poolA.BondedTokens)))

	// specific case from random tests
	poolTokens := types.NewDec(5102)
	delShares := types.NewDec(115)
	validator := Validator{
		Owner:           addr1,
		PubKey:          pk1,
		Status:          types.Bonded,
		Tokens:          poolTokens,
		DelegatorShares: delShares,
	}
	pool := Pool{
		BondedTokens: types.NewDec(248305),
		LooseTokens:  types.NewDec(232147),
	}
	shares := types.NewDec(29)
	_, newPool, tokens := validator.RemoveDelShares(pool, shares)

	exp, err := types.NewDecFromStr("1286.5913043477")
	require.NoError(t, err)

	require.True(types.DecEq(t, exp, tokens))

	require.True(types.DecEq(t,
		newPool.LooseTokens.Add(newPool.BondedTokens),
		pool.LooseTokens.Add(pool.BondedTokens)))
}

func TestUpdateStatus(t *testing.T) {
	pool := InitialPool()
	pool.LooseTokens = types.NewDec(100)

	validator := NewValidator(addr1, pk1, Description{})
	validator, pool, _ = validator.AddTokensFromDel(pool, types.NewInt(100))
	require.Equal(t, types.Unbonded, validator.Status)
	require.Equal(t, int64(100), validator.Tokens.RoundInt64())
	require.Equal(t, int64(0), pool.BondedTokens.RoundInt64())
	require.Equal(t, int64(100), pool.LooseTokens.RoundInt64())

	validator, pool = validator.UpdateStatus(pool, types.Bonded)
	require.Equal(t, types.Bonded, validator.Status)
	require.Equal(t, int64(100), validator.Tokens.RoundInt64())
	require.Equal(t, int64(100), pool.BondedTokens.RoundInt64())
	require.Equal(t, int64(0), pool.LooseTokens.RoundInt64())

	validator, pool = validator.UpdateStatus(pool, types.Unbonding)
	require.Equal(t, types.Unbonding, validator.Status)
	require.Equal(t, int64(100), validator.Tokens.RoundInt64())
	require.Equal(t, int64(0), pool.BondedTokens.RoundInt64())
	require.Equal(t, int64(100), pool.LooseTokens.RoundInt64())
}

func TestPossibleOverflow(t *testing.T) {
	poolTokens := types.NewDec(2159)
	delShares := types.NewDec(391432570689183511).Quo(types.NewDec(40113011844664))
	validator := Validator{
		Owner:           addr1,
		PubKey:          pk1,
		Status:          types.Bonded,
		Tokens:          poolTokens,
		DelegatorShares: delShares,
	}
	pool := Pool{
		LooseTokens:  types.NewDec(100),
		BondedTokens: poolTokens,
	}
	tokens := int64(71)
	msg := fmt.Sprintf("validator %#v", validator)
	newValidator, _, _ := validator.AddTokensFromDel(pool, types.NewInt(tokens))

	msg = fmt.Sprintf("Added %d tokens to %s", tokens, msg)
	require.False(t, newValidator.DelegatorShareExRate().LT(types.ZeroDec()),
		"Applying operation \"%s\" resulted in negative DelegatorShareExRate(): %v",
		msg, newValidator.DelegatorShareExRate())
}

func TestHumanReadableString(t *testing.T) {
	validator := NewValidator(addr1, pk1, Description{})

	// NOTE: Being that the validator's keypair is random, we cannot test the
	// actual contents of the string.
	valStr, err := validator.HumanReadableString()
	require.Nil(t, err)
	require.NotEmpty(t, valStr)
}

func TestValidatorMarshalUnmarshalJSON(t *testing.T) {
	validator := NewValidator(addr1, pk1, Description{})
	js, err := codec.Cdc.MarshalJSON(validator)
	require.NoError(t, err)
	require.NotEmpty(t, js)
	require.Contains(t, string(js), "\"consensus_pubkey\":\"cosmosvalconspu")
	got := &Validator{}
	err = codec.Cdc.UnmarshalJSON(js, got)
	assert.NoError(t, err)
	assert.Equal(t, validator, *got)
}

/*func TestValidatorSetInitialCommission(t *testing.T) {
	val := NewValidator(addr1, pk1, Description{})
	testCases := []struct {
		validator   Validator
		commission  Commission
		expectedErr bool
	}{
		{val, NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()), false},
		{val, NewCommission(sdk.ZeroDec(), types.NewDecWithPrec(-1, 1), sdk.ZeroDec()), true},
		{val, NewCommission(sdk.ZeroDec(), types.NewDec(15000000000), sdk.ZeroDec()), true},
		{val, NewCommission(types.NewDecWithPrec(-1, 1), sdk.ZeroDec(), sdk.ZeroDec()), true},
		{val, NewCommission(types.NewDecWithPrec(2, 1), types.NewDecWithPrec(1, 1), sdk.ZeroDec()), true},
		{val, NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), types.NewDecWithPrec(-1, 1)), true},
		{val, NewCommission(sdk.ZeroDec(), types.NewDecWithPrec(1, 1), types.NewDecWithPrec(2, 1)), true},
	}

	for i, tc := range testCases {
		val, err := tc.validator.SetInitialCommission(tc.commission)

		if tc.expectedErr {
			require.Error(t, err,
				"expected error for test case #%d with commission: %s", i, tc.commission,
			)
		} else {
			require.NoError(t, err,
				"unexpected error for test case #%d with commission: %s", i, tc.commission,
			)
			require.Equal(t, tc.commission, val.Commission,
				"invalid validator commission for test case #%d with commission: %s", i, tc.commission,
			)
		}
	}
} */
