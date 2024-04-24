#!/bin/bash

cd "$(dirname "$0")"
cd ../../

echo "Doing work in directory $PWD"

BASE_DIR="$PWD"
BRANCH="$(git rev-parse --abbrev-ref HEAD)"
LATEST_SMR_COMMIT="$(git rev-parse --short $BRANCH)"

cd "$BASE_DIR"

echo "***********************************************"
echo "$BASE_DIR/../implementations/$DIRNAME"
echo "***********************************************"

go build -ldflags '-s -w' -o client
chmod +x client
upx -9 -k client