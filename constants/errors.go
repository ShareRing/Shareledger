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
