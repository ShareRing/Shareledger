#!/usr/bin/env bash

mockgen_cmd="go run github.com/golang/mock/mockgen"

$mockgen_cmd -source=x/document/types/expected_keepers.go -package testutil -destination x/document/testutil/expected_keepers_mocks.go
