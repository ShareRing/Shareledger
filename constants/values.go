package constants

// APP ACCOUNT
const DEFAULT_DENOM = "SHR"
const DEFAULT_AMOUNT = 0
const PREFIX_ADDRESS = "account:" // address to string to store in Auth Module

// STORE
const STORE_BANK = "bank"
const STORE_BOOKING = "booking"
const STORE_ASSET = "asset"
const STORE_AUTH = "auth"
const STORE_POS = "pos"
const STORE_EXCHANGE = "excrate"

// MESSAGE TYPE
const MESSAGE_AUTH = "auth"
const MESSAGE_ASSET = "asset"
const MESSAGE_BANK = "bank"
const MESSAGE_BOOKING = "booking"
const MESSAGE_POS = "pos"
const MESSAGE_EXCHANGE_RATE = "exchangerate"

// ALLOWED DENOM
var DENOM_LIST = map[string]bool{"SHRP": true, "SHR": true}
var ALL_DENOMS = []string{"SHRP", "SHR"}
var BOOKING_DENOM = "SHRP"
var POS_DENOM = "SHR"
var EXCHANGABLE_FEE_DENOM = "SHRP"
var DEFAULT_RESERVE = "405C725BC461DCA455B8AA84769E8ACE6B3763F4"
var POS_BLOCK_REWARD = int64(5)

//POS Constant
var MIN_MASTER_NODE_TOKEN int64 = 2000000

// EXCHANGE
var RESERVE_ACCOUNTS = []string{
	"405C725BC461DCA455B8AA84769E8ACE6B3763F4",
	"B87D5A84F7DCE488BA2FCBDD2057023561BC05A4",
}
