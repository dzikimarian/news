Compile for linux

    docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.17 go build -v

Compile for windows

    docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=windows -e GOARCH=amd64 golang:1.17 go build -v

Compile for OSX

    docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=darwin -e GOARCH=amd64 golang:1.17 go build -v