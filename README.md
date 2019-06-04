# fortune-cookie

A fortune cookie generator (based on Go, GraphQL)

```bash
$ go version
go version go1.12.5 linux/amd64

$ export GOPATH=${HOME}/go
$ export GOBIN=${GOPATH}/bin
$ export PATH=${PATH}:${GOBIN}

$ go build ./...

$ go install
$ fortune-cookie # grab your cookie (while it's hot)

{"data":{"tellMe":"In skating over thin ice our safety is in our speed. Ralph Emerson"}} 

```