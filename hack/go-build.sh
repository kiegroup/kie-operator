#!/bin/sh

IMAGE=quay.io/kiegroup/kie-cloud-operator
TAG=v0.1.0

go generate ./...
if [[ -z ${CI} ]]; then
    source hack/go-test.sh
    operator-sdk build ${IMAGE}:${TAG}
else
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o build/_output/bin/kie-cloud-operator github.com/kiegroup/kie-cloud-operator/cmd/manager
fi