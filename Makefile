all:
	go build sctp-shell.go

release:
	go build -ldflags="-w -s" sctp-shell.go
