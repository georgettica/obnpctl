#!/bin/bash

set -eEuo pipefail

CURRENT_DIR="$(dirname "$0")"

./bin/kustomize build "${CURRENT_DIR}" | oc get -f - --ignore-not-found

if [[ $(./bin/kustomize build "${CURRENT_DIR}" | ./bin/kubectl-slice -f - --include-kind=Deployment --stdout | oc get -f - --ignore-not-found | wc -l ) -eq 0 ]]; then
	echo "there are no Deployemnts, no need to wait for them to be deleted"
	exit
fi
./bin/kustomize build "${CURRENT_DIR}" | ./bin/kubectl-slice -f - --include-kind=Deployment --stdout | oc wait  -f - --for=delete
