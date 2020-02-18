all: build/osx-amd64/http-to-quintype-proxy build/linux-386/http-to-quintype-proxy build/linux-amd64/http-to-quintype-proxy

build/osx-amd64/http-to-quintype-proxy: main.go
	GOOS=darwin GOARCH=amd64 go build -o $@

build/linux-386/http-to-quintype-proxy: main.go
	GOOS=linux GOARCH=386 go build -o $@

build/linux-amd64/http-to-quintype-proxy: main.go
	GOOS=linux GOARCH=amd64 go build -o $@
