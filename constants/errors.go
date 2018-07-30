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