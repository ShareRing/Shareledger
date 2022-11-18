package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		WasmMasterBuilderPercent: 12.5,
		WasmContractAdminPercent: 12.5,
		WasmDevelopmentPercent:   25,
		WasmValidatorPercent:     50,

		NativeValidatorPercent:   50,
		NativeDevelopmentPercent: 50,

		BuilderWindows: 1000,
		TxThreshold:    1000,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	// total wasm percent is 100
	totalWasm := p.WasmMasterBuilderPercent + p.WasmContractAdminPercent + p.WasmDevelopmentPercent + p.WasmValidatorPercent
	if totalWasm != 100 {
		return ErrInvalidParams.Wrapf("total wasm percent is: %v, expected: 100", totalWasm)
	}
	// total native percent is 100
	totalNative := p.NativeDevelopmentPercent + p.NativeValidatorPercent
	if totalNative != 100 {
		return ErrInvalidParams.Wrapf("total native percent is: %v, expected 100", totalNative)
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
