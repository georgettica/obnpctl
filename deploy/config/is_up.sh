#!/bin/bash

set -eEuo pipefail

CURRENT_DIR="$(dirname "$0")"

./bin/kustomize build "${CURRENT_DIR}" | ./bin/kubectl-slice -f - --include-kind=PostgresCluster --stdout | oc wait  -f - --for=condition=PGBackRestRepoHostReady
