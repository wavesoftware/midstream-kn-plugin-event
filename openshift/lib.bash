#!/usr/bin/env bash

set -Eeuo pipefail

# go.get-tool will 'go get' any package $2 and install it to $1. Replaces, can
# be set with $3.
function go.get-tool {
  local tmp_dir replace
  if ! [ -f "${1}" ]; then
    replace="${3:-}"
    tmp_dir="$(mktemp -d)"
    trap 'rm -rf ${tmp_dir}' EXIT
    cd "$tmp_dir"
    go mod init tmp
    if [ -n "${replace:-}" ]; then
      go mod edit -replace "${replace}"
    fi
    echo "Downloading ${2}"
    GOBIN="$(dirname "${1}")" go get "${2}"
  fi
}