package testutil

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func UnMarshalJson(t *testing.T, outPut []byte, dataPlaceHolder interface{}) {
	if reflect.TypeOf(dataPlaceHolder).Kind() != reflect.Ptr {
		t.Logf("the struct type must be pointer")
		t.Fail()
	}
	err := json.Unmarshal(outPut, dataPlaceHolder)
	require.NoError(t, err)
}
