build:
	go build -o build/linux/RefInator cmd/main.go

mac_build:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o build/mac/RefInator cmd/main.go

run_build:
	go build -o build/linux/RefInator cmd/main.go
	build/linux/RefInator