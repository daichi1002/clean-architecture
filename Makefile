lint:
	golangci-lint run

run:
	go run main.go

test:
	go test ./... -v -coverprofile=profile.out -p 1
	go tool cover -func=profile.out
