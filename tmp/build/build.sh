#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

if ! which go > /dev/null; then
	echo "golang needs to be installed"
	exit 1
fi

BIN_DIR="$(pwd)/tmp/_output/bin"
mkdir -p ${BIN_DIR}
PROJECT_NAME="kie-cloud-operator"
REPO_PATH="github.com/kiegroup/kie-cloud-operator"
BUILD_PATH="${REPO_PATH}/cmd/${PROJECT_NAME}"
echo "Building "${PROJECT_NAME}"..."
go generate ${REPO_PATH}/...
go fmt ${REPO_PATH}/...
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BIN_DIR}/${PROJECT_NAME} $BUILD_PATH
