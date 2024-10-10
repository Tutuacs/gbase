build:
	@GOOS=windows GOARCH=amd64 go build -o bin/gbase.exe main.go
