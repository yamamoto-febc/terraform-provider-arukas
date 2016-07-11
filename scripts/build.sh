#!/bin/bash

set -e

OS="darwin linux windows"
ARCH="amd64 386"

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
        GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 govendor build -o $binary builtin/bins/provider-arukas/main.go
        zip -r "bin/terraform-provider-arukas_$arch" $binary
        rm -f $binary
    done
done
