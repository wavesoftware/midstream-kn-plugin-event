#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(realpath "$(dirname "${BASH_SOURCE[0]:-$0}")/../..")"
# shellcheck source=../lib.bash
source "${project_dir}/openshift/lib.bash"

go.run-tool \
  github.com/cardil/deviate/cmd/deviate@v0.1.2 \
  sync --config "${project_dir}/.deviate.yaml"
