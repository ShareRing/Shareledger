package pos

const (
	EventTypeValidatorCreated     = "ValidatorCreated"
	EventTypeEditValidator        = "EditValidator"
	EventTypeDelegated            = "Delegated"
	EventTypeBeginUnbonding       = "BeginUnbonding"
	EventTypeCompleteUnbonding    = "CompleteUnbonding"
	EventTypeWithdraw             = "Withdraw"
	EventTypeBeginRedelegation    = "BeginRedelegation"
	EventTypeCompleteRedelegation = "CompleteRedelegation"

	AttributeValidator    = "Validator"
	AttributeDelegator    = "Delegator"
	AttributeSrcValidator = "SrcValidator"
	AttributeDstValidator = "DstValidator"
	AttributeMoniker      = "Moniker"
	AttributeIdentity     = "Identity"
)
