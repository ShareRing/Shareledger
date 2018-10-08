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

// MESSAGE TYPE
const MESSAGE_AUTH = "auth"
const MESSAGE_ASSET = "asset"
const MESSAGE_BANK = "bank"
const MESSAGE_BOOKING = "booking"
const MESSAGE_POS = "pos"

// ALLOWED DENOM
var DENOM_LIST = map[string]bool{"SHRP": true, "SHR": true}
var BOOKING_DENOM = "SHRP"
