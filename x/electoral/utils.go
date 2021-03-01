package electoral

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func GenBookID(inp interface{}) (string, error) {
	h := sha256.New()

	enc, err := json.Marshal(inp)
	if err != nil {
		return "", err
	}

	h.Write(enc)
	hash := h.Sum(nil)

	return hex.EncodeToString(hash)[:20], nil
}
