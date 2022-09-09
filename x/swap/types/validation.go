package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

func (m *QuerySwapRequest) BasicValidate() error {
	if !SwapStatusSupported(m.Status) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s status is not supported", m.Status)
	}
	return nil
}
