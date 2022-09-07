cover: go-test-cover go-tool-cover

go-test-cover:
	go test ./... -cover -v -coverprofile=cover.out

go-tool-cover:
	go tool cover -html cover.out
