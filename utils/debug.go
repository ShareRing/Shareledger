package utils

import "fmt"

func Format(stringFormat string, params... interface{}) string {
	return fmt.Sprintf(stringFormat, params...)
}