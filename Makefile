local:
	go run app/main.go local

lint:
	golangci-lint run
	