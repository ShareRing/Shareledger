package tags

var (
	//Key - String type

	Event        = "Event"
	Validator    = "Validator"
	Delegator    = "Delegator"
	SrcValidator = "SrcValidator"
	DstValidator = "DstValidator"
	Moniker      = "Moniker"
	Identity     = "Identity"

	//Value -  []byte

	ValidatorCreated     = "ValidatorCreated"
	EditValidator        = "EditValidator"
	Delegated            = "Delegated"
	BeginUnbonding       = "BeginUnbonding"
	CompleteUnbonding    = "CompleteUnbonding"
	Withdraw             = "Witdrawed"
	BeginRedelegation    = "BeginRedelegation"
	CompleteRedelegation = "CompleteRedelegation"
)
