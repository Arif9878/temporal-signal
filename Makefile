worker-start:
	go run ./worker/main.go

worker-signal:
	go run ./start/main.go

test:
	go test -v ./...