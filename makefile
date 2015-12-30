default:
	GOOS=freebsd GOARCH=amd64 go run ./main.go

darwin:
	go run ./main.go
