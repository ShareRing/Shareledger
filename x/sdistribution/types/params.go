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
		WasmMasterBuilder: .125,
		WasmContractAdmin: .125,
		WasmDevelopment:   .25,
		WasmValidator:     .5,

		NativeValidator:   .5,
		NativeDevelopment: .5,

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
	// total wasm is 1
	totalWasm := p.WasmMasterBuilder + p.WasmContractAdmin + p.WasmDevelopment + p.WasmValidator
	if totalWasm != 1 {
		return ErrInvalidParams.Wrapf("total wasm is: %v, expected: 1", totalWasm)
	}
	// total native is 1
	totalNative := p.NativeDevelopment + p.NativeValidator
	if totalNative != 1 {
		return ErrInvalidParams.Wrapf("total native is: %v, expected 1", totalNative)
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
