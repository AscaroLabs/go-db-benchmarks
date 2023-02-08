build: bs bp bt

bs:
	GOOS=linux GOARCH=amd64 go build -a -tags musl -a -installsuffix cgo -o up cmd/setup/setup.go
bp:
	GOOS=linux GOARCH=amd64 go build -a -tags musl -a -installsuffix cgo -o postgres cmd/postgres/main.go
bt:
	GOOS=linux GOARCH=amd64 go build -a -tags musl -a -installsuffix cgo -o tnt cmd/tarantool/main.go

