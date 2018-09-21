package requests

const (
	NodeUrl     = "http://192.168.1.234:56657"
	// NodeUrl 	= "http://118.69.238.23:5667"
	StatusUri   = "/status"
	QueryUri    = "/abci_query?path=\"app/query\"&data=%s"
	BroadcastTx = "/broadcast_tx_commit?tx=%s"
)
