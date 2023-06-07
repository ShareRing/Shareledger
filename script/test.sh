#!/usr/bin/env bash

# Note that you should run this script from root dir
# sh ./script/test.sh

echo "============START TO RUN TESTSUITE===================="
module_list=(id swap asset booking document distributionx electoral)
pids=""

mkdir -p coverage

go test ./x/... -coverprofile coverage.out &
pids+=" $!"

for i in "${!module_list[@]}"; do
	go test --tags e2e -coverprofile=coverage/${module_list[$i]}.out -coverpkg=./... ./tests/e2e/${module_list[$i]} &
  pids+=" $!"
  module_list[$i]=./coverage/${module_list[$i]}.out
done
go install github.com/wadey/gocovmerge@latest

# make sure that all tests is pass
for p in $pids; do
	if ! wait $p; then
    echo "============RUN THE TEST FAILED!===================="
    # kill all child process and exit
    pkill -P $$
		exit 1
	fi
done

gocovmerge ${module_list[@]} ./coverage.out >./coverage/merged_coverage.out

# remove generated and simulation file
grep -vE '(pb(\.gw)?\.go)|simulation|pulsar' ./coverage/merged_coverage.out >./coverage/coverage.out

go tool cover -func ./coverage/coverage.out | tail -n1

echo "============RUN THE TEST SUCCESSFUL===================="
