package types

import (
	"github.com/google/uuid"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateIds_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateIds
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateIds{
				IssuerAddress: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateIds{
				IssuerAddress: sample.AccAddress(),
				BackupAddress: []string{sample.AccAddress()},
				ExtraData:     []string{"data"},
				Id:            []string{uuid.New().String()},
				OwnerAddress:  []string{sample.AccAddress()},
			},
		},
		{
			name: "valid data",
			msg: MsgCreateIds{
				IssuerAddress: "shareledger1l3pg3zd0u5p3v5wfqavh0gsr83zep8kv5y900z",
				BackupAddress: []string{
					"shareledger1zf6q3twxs9dgw0dhjz0msve5ez3638vgddvgar",
					"shareledger1k57s3hqnky2pawny6lx0j8xnnjtshtcnpe6ewu",
					"shareledger17vlzwgh6k7y3trday57fw6c4nrv6gnmfmj25qa",
				},
				OwnerAddress: []string{
					"shareledger1zf6q3twxs9dgw0dhjz0msve5ez3638vgddvgar",
					"shareledger1k57s3hqnky2pawny6lx0j8xnnjtshtcnpe6ewu",
					"shareledger17vlzwgh6k7y3trday57fw6c4nrv6gnmfmj25qa",
				},
				Id: []string{
					"id-1",
					"id-2",
					"id-3",
				},
				ExtraData: []string{
					"hello1",
					"hello2",
					"hello3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
