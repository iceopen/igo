OPTS=""

#linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $OPTS -ldflags "$X" -o bin/igo && tar zcfv "bin/igo-linux-amd64.tar.gz" bin/igo

#darwin
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $OPTS -ldflags "$X" -o bin/igo && tar zcfv "bin/igo-darwin-amd64.tar.gz" bin/igo

#windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $OPTS  -ldflags "$X" -o bin/igo.exe && tar zcfv "bin/igo-windows-amd64.tar.gz" bin/igo
