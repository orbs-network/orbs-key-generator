#!/bin/sh
rm -rf ./_bin

mkdir -p ./_bin

VERSION="v0.2.0"

echo "\n\n*** MAC:"
GOOS=darwin GOARCH=amd64 go build -o _bin/orbs-key-generator
tar -zcvf ./_bin/orbskeygenerator-mac-$VERSION.tar.gz ./_bin/orbs-key-generator
rm ./_bin/orbs-key-generator

cd ./_bin

openssl sha256 orbskeygenerator-mac-$VERSION.tar.gz >> ./checksums.txt
