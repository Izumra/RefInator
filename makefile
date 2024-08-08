build:
	go build -o build/linux/RefInator cmd/main.go

run_build:
	go build -o build/linux/RefInator cmd/main.go
	build/linux/RefInator