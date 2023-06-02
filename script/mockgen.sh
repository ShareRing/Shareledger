#!/usr/bin/env bash

mockgen_cmd="go run github.com/golang/mock/mockgen"

$mockgen_cmd -source=x/document/types/expected_keepers.go -package testutil -destination x/document/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=x/booking/types/expected_keepers.go -package testutil -destination x/booking/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=x/electoral/types/expected_keepers.go -package testutil -destination x/electoral/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=x/electoral/keeper/dependency.go -package testutil -destination x/electoral/testutil/dependency_mocks.go