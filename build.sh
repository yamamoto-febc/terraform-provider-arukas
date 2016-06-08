#!/bin/bash

set -e

OS="darwin linux windows"
ARCH="amd64 386"

echo "Ensuring code quality"
#go vet ./...
gofmt -w .

ver=$(cd $GOPATH/src/github.com/hashicorp/terraform && git describe --abbrev=0 --tags)
echo "VERSION terraform '$ver'"

rm -Rf bin/
mkdir bin/

for GOOS in $OS; do
    for GOARCH in $ARCH; do
        arch="$GOOS-$GOARCH"
        binary="terraform-provider-arukas"
        if [ "$GOOS" = "windows" ]; then
          binary="${binary}.exe"
        fi
        echo "Building $binary $arch"
        GOOS=$GOOS GOARCH=$GOARCH godep go build -o $binary builtin/bins/provider-arukas/main.go
        zip -r "bin/terraform-provider-arukas_$arch" $binary
        rm -f $binary
    done
done
