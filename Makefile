build:
	go build -o build/shareledger -mod=readonly cmd/shareledger/main.go

build_linux_arm64:
	env GOOS=linux GOARCH=arm64 go build -o build/shareledger_linux_arm64 cmd/shareledger/main.go

build_linux_amd64:
	env GOOS=linux GOARCH=arm64 go build -o build/shareledger_linux_arm64 cmd/shareledger/main.go

build_windows_amd64:
	env GOOS=windows GOARCH=amd64 go build -o build/shareledger_windows_arm64.exe cmd/shareledger/main.go


build_all: build build_linux_arm64 build_linux_amd64 build_windows_amd64

.PHONY: build build_linux_arm64 build_linux_amd64 build_windows_amd64 build_all
