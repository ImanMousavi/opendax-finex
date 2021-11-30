serve:
	go mod download
	go generate
	go get ./...
	go run -mod=mod ./cmd/server

clean:
	git clean -qfdX
