package constants

// ENCODING ERROR
const ERROR_ENCODING = "Encoding %s failed."
const ERROR_DECODING = "Decoding %s failed."

// STORE OPERATION ERRORS
const ERROR_STORE_RETRIEVAL = "Retrieval %s from Store %s failed."
const ERROR_STORE_UPDATE = "Update %s to Store %s failed."
const ERROR_STORE_NOT_FOUND = "%s cannot be found in Store %s."

// NOT FOUND
const BOOKING_MISMATCH_RENTER = "Mismatch renter from Booking request %s with Complete request %s"
const BOOKING_COMPLETED_ERROR = "The booking %s is already completed."
const BOOKING_ASSET_NOT_RENTED = "Asset %s is not rented."
const BOOKING_ASSET_RENTED = "Asset %s is already rented."
const BOOKING_INSUFFICIENT_BALANCE = "Account %s has insuficient balance."

// SHRAccount
const SHRACCOUNT_EXISITNG_ADDRESS = "Address already exists."
const SHRACCOUNT_INVALID_ADDRESS = "Invalid address."

// Proto Error
const ACCOUNT_INVALID_STRUCT = "accountMapper requires a struct proto BaseAccount, or a pointer to one"
const ACCOUNT_INVALID_INTERFACE = "accountMapper requries a proto BaseAccount, but %v doesn't implement BaseAccount interface."

// Tx Fee Calculation
const INSUFFICIENT_BALANCE = "Account %s has insufficient balance."
const INVALID_TX_FEE = "Invalid transaction fee %s."

// Two separators found
const DEC_TWO_SEPARATORS = "Two separators found at %d and %d."
const DEC_INVALID_DECIMALS = "Too many decimal digits in fractional part. %s"

// POS
const POS_WITHDRAWAL_ERROR = "Error in withdrawl for delegator %s."
const POS_DELEGATION_NOT_FOUND = "Delegation with address %X not found."
const POS_VALIDATOR_DIST_NOT_FOUND = "Validator Distribution Info for %X not found."
const POS_INVALID_VALIDATOR_ADDRESS = "Validator Address is in incorrect format. Error: %s"
const POS_MARSHAL_ERROR = "Marshal to JSON failed. %s"
const POS_INVALID_PARAMS = "Unmarshal to Query params failed. %s"

// Exchange
const EXC_INVALID_DENOM = "Invalid Denom. Required %s. Provided %s."
const EXC_JSON_MARSHAL = "MarshalJSON failed. %s"
const EXC_EXCHANGE_RATE_NOT_FOUND = "Exchange Rate from %s to %s not found."
const EXC_SAME_DENOM = "FromDenom and ToDenom should be different. Not the same %s."
const EXC_INVALID_RATE = "Rate must be larger than 0. Provided rate %s."
const EXC_INVALID_AMOUNT = "Amount must be larger than 0. Provided amount %s."
const EXC_INVALID_RESERVE = "Invalid Reserve %s."
const EXC_INSUFFICIENT_BALANCE = "Account (%s < %s) or Reserve (%s < %s) has insufficient amount."
const EXC_ALREADY_EXIST = "Exchange Rate from %s to %s has already existed."

// RESERVE
const RES_RESERVE_ONLY = "Only priviledged accounts can execute this transaction."
const RES_OWN_ACCOUNT = "An account can only burn Coins of its own. Account %s != Signer %s."

// BANK
const BANK_INVALID_BURNT_DENOM = "Only booking denom %s is allowed to be burnt."
