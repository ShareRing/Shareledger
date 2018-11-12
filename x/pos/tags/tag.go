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

	ValidatorCreated     = []byte("ValidatorCreated")
	Delegated            = []byte("Delegated")
	BeginUnbonding       = []byte("BeginUnbonding")
	CompleteUnbonding    = []byte("CompleteUnbonding")
	Withdraw             = []byte("Witdrawed")
	BeginRedelegation    = []byte("BeginRedelegation")
	CompleteRedelegation = []byte("CompleteRedelegation")
)
