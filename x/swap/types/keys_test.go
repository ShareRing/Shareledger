package types

import "testing"

func TestPastTxEventKeyReverser(t *testing.T) {
	inputHash := "hash132323232"
	var inputLogIndex uint64 = 34
	inPutKey := PastTxEventKey(inputHash, inputLogIndex)

	txHash, logIndex, err := PastTxEventKeyReverser(inPutKey)

	if err != nil {
		t.Errorf("revese key fail %s", err)
	}
	if txHash != inputHash {
		t.Errorf("txHash not equal. got=%s expect= %s", txHash, inputHash)
	}
	if logIndex != inputLogIndex {
		t.Errorf("logIndex not equal. got=%d expect= %d", logIndex, inputLogIndex)
	}
}
