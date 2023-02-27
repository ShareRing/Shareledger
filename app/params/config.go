package params

import (
	"strings"

	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// BypassMinFeeMsgTypesKey defines the configuration key for the
	// BypassMinFeeMsgTypes value.
	//nolint: gosec
	BypassMinFeeMsgTypesKey = "bypass-min-fee-msg-types"

	// customConfigTemplate defines custom application configuration TOML template.
	customConfigTemplate = `
###############################################################################
###                        Custom Configuration                        ###
###############################################################################
# bypass-min-fee-msg-types defines custom message types the operator may set that
# will bypass minimum fee checks during CheckTx.
#
# Example:
# ["/ibc.core.channel.v1.MsgRecvPacket", "/ibc.core.channel.v1.MsgAcknowledgement", ...]
bypass-min-fee-msg-types = [{{ range .BypassMinFeeMsgTypes }}{{ printf "%q, " . }}{{end}}]
`
)

const (
	Bech32PrefixAccAddr = "shareledger"
)

var (
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key.
	Bech32PrefixAccPub = Bech32PrefixAccAddr + "pub"
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address.
	Bech32PrefixValAddr = Bech32PrefixAccAddr + "valoper"
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key.
	Bech32PrefixValPub = Bech32PrefixAccAddr + "valoperpub"
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address.
	Bech32PrefixConsAddr = Bech32PrefixAccAddr + "valcons"
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key.
	Bech32PrefixConsPub = Bech32PrefixAccAddr + "valconspub"
)

// CustomConfigTemplate defines custom application configuration TOML
// template. It extends the core SDK template.
func CustomConfigTemplate() string {
	config := serverconfig.DefaultConfigTemplate
	lines := strings.Split(config, "\n")
	// add the config at the second line of the file
	lines[2] += customConfigTemplate
	return strings.Join(lines, "\n")
}

// CustomAppConfig defines custom application configuration.
type CustomAppConfig struct {
	serverconfig.Config

	// BypassMinFeeMsgTypes defines custom message types the operator may set that
	// will bypass minimum fee checks during CheckTx.
	BypassMinFeeMsgTypes []string `mapstructure:"bypass-min-fee-msg-types"`
}

func SetAddressPrefixes() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
}
