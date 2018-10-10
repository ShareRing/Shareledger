package constants

// Level of each message types
type FeeLevel int

const (
	HIGH FeeLevel = 0
	MED  FeeLevel = 1
	LOW  FeeLevel = 2
)

var LEVELS = map[string]FeeLevel{
	"MsgSend":     LOW,
	"MsgCreate":   HIGH,
	"MsgUpdate":   MED,
	"MsgDelete":   LOW,
	"MsgBook":     HIGH,
	"MsgComplete": MED,
}

var FEE_LEVELS = map[FeeLevel]int{
	HIGH: 3,
	MED:  2,
	LOW:  1,
}

const FEE_DENOM = "SHR"
