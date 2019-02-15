package tags

var (
	//Key - String type
	FromAddress    = "FromAddress"
	ToAddress      = "ToAddress"
	Amount         = "Amount"
	Event          = "Event"
	AccountAddress = "AccountAddress"

	//Value -  []byte
	Transfered = "Transfered" //Transfer event fromAddress To Address
	Credit     = "Credit"     //event for credit
)
