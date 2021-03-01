package types

import (
	"fmt"
	"strings"
)

const (
	IdSignerActive          = "active"
	IdSignerInactive        = "inactive"
	defaultSHRPLoaderStatus = IdSignerInactive
)

type IdSigner struct {
	Status string `json:"status"`
}

func NewIdSigner() IdSigner {
	return IdSigner{
		Status: defaultSHRPLoaderStatus,
	}
}

func (il IdSigner) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Status %s`, il.Status))
}
