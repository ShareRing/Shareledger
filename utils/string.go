package utils

import "fmt"

func ByteToString(inp []byte) string {
	return  fmt.Sprintf("%x", inp)
}