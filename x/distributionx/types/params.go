package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// Parameter store keys
var (
	WasmMasterBuilderKey = []byte("wasmmasterbuilder")
	WasmContractAdminKey = []byte("wasmcontractadmin")
	WasmDevelopmentKey   = []byte("wasmdevelopment")
	WasmValidatorKey     = []byte("wasmvalidator")
	NativeValidatorKey   = []byte("nativevalidator")
	NativeDevelopmentKey = []byte("nativedevelopment")
	BuilderWindowsKey    = []byte("builderwindows")
	TxThresholdKey       = []byte("txthreshold")
	DevPoolAccountKey    = []byte("devpoolaccount")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		WasmMasterBuilder: sdk.NewDecWithPrec(125, 3),
		WasmContractAdmin: sdk.NewDecWithPrec(125, 3),
		WasmDevelopment:   sdk.NewDecWithPrec(250, 3),
		WasmValidator:     sdk.NewDecWithPrec(500, 3),
		NativeValidator:   sdk.NewDecWithPrec(500, 3),
		NativeDevelopment: sdk.NewDecWithPrec(500, 3),

		BuilderWindows: 1000,
		TxThreshold:    100,
		DevPoolAccount: "",
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
// TODO: add validator for each field
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(WasmMasterBuilderKey, &p.WasmMasterBuilder, noValidate),
		paramtypes.NewParamSetPair(WasmContractAdminKey, &p.WasmContractAdmin, noValidate),
		paramtypes.NewParamSetPair(WasmDevelopmentKey, &p.WasmDevelopment, noValidate),
		paramtypes.NewParamSetPair(WasmValidatorKey, &p.WasmValidator, noValidate),
		paramtypes.NewParamSetPair(NativeValidatorKey, &p.NativeValidator, noValidate),
		paramtypes.NewParamSetPair(NativeDevelopmentKey, &p.NativeDevelopment, noValidate),
		paramtypes.NewParamSetPair(BuilderWindowsKey, &p.BuilderWindows, noValidate),
		paramtypes.NewParamSetPair(TxThresholdKey, &p.TxThreshold, noValidate),
		paramtypes.NewParamSetPair(DevPoolAccountKey, &p.DevPoolAccount, noValidate),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	// total wasm is 1
	if !p.WasmMasterBuilder.Add(p.WasmContractAdmin).Add(p.WasmDevelopment).Add(p.WasmValidator).Equal(sdk.NewDec(1)) {
		return ErrInvalidParams.Wrapf("total wasm is not equal: 1")
	}
	// total native is 1
	if !p.NativeDevelopment.Add(p.NativeValidator).Equal(sdk.NewDec(1)) {
		return ErrInvalidParams.Wrapf("total native is not equal 1")
	}

	if p.TxThreshold == 0 {
		return ErrInvalidParams.Wrapf("invalid TxThreshold: %d", p.TxThreshold)
	}
	err := sdk.VerifyAddressFormat([]byte(p.DevPoolAccount))
	if err != nil {
		return ErrInvalidParams.Wrapf("invalid DevPoolAccount :%s", err)
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func noValidate(i interface{}) error {
	return nil
}
