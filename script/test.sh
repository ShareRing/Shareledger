#!/usr/bin/env bash

# Note that you should run this script from root dir
# sh ./script/test.sh

module_list=(id swap asset booking document distributionx electoral)


go test ./x/...

mkdir -p coverage

go test ./x/... -coverprofile coverage.out

for i in "${!module_list[@]}"; do
  go test --tags e2e -coverprofile=coverage/${module_list[$i]}.out -coverpkg=./... ./tests/e2e/${module_list[$i]}
  module_list[$i]=./coverage/${module_list[$i]}.out
done
go install github.com/wadey/gocovmerge@latest

gocovmerge ${module_list[@]} ./coverage.out > ./coverage/after_merge_coverate.out




# todo: wait all test done merge coverage
# coverage.out

# remove .pb.go gw.pb.go
grep -vE 'pb(\.gw)?\.go' ./coverage/after_merge_coverate.out | grep -v simulation | grep -v pulsar > ./coverage/coverage.out