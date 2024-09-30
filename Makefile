compile:
	GOOS=linux GOARCH=amd64 go build -o ./bin/boop_linux_amd64 ./cmd/boop # linux
	GOOS=darwin GOARCH=amd64 go build -o ./bin/boop_darwin_amd64 ./cmd/boop # macOS
	GOOS=darwin GOARCH=arm64 go build -o ./bin/boop_darwin_arm64 ./cmd/boop # Apple M1
	GOOS=windows GOARCH=amd64 go build -o ./bin/boop_win_amd64.exe ./cmd/boop # Windows