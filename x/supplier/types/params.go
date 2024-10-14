package types

import (
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/pokt-network/poktroll/app/volatile"
)

var (
	_ paramtypes.ParamSet = (*Params)(nil)

	KeyMinStake   = []byte("MinStake")
	ParamMinStake = "min_stake"
	// TODO_MAINNET: Determine the default value.
	DefaultMinStake = cosmostypes.NewInt64Coin("upokt", 1000000) // 1 POKT
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(minStake *cosmostypes.Coin) Params {
	return Params{
		MinStake: minStake,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(&DefaultMinStake)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			KeyMinStake,
			&p.MinStake,
			ValidateMinStake,
		),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := ValidateMinStake(p.MinStake); err != nil {
		return err
	}

	return nil
}

// ValidateMinStake validates the MinStake param.
func ValidateMinStake(minStakeAny any) error {
	minStakeCoin, ok := minStakeAny.(*cosmostypes.Coin)
	if !ok {
		return ErrSupplierParamInvalid.Wrapf("invalid parameter type: %T", minStakeAny)
	}

	if minStakeCoin == nil {
		return ErrSupplierParamInvalid.Wrapf("missing min_stake")
	}

	if minStakeCoin.Denom != volatile.DenomuPOKT {
		return ErrSupplierParamInvalid.Wrapf(
			"invalid min_stake denom %q; expected %q",
			minStakeCoin.Denom, volatile.DenomuPOKT,
		)
	}

	if minStakeCoin.IsZero() || minStakeCoin.IsNegative() {
		return ErrSupplierParamInvalid.Wrapf("min_stake amount must be greater than 0: got %s", minStakeCoin)
	}

	return nil
}
