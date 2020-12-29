all:
	go build sctp-shell.go

relase:
	go build -ldflags="-w -s" sctp-shell.go
