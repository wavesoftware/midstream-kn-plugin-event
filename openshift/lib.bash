#!/usr/bin/env bash

set -Eeuo pipefail

# go.get-tool will 'go get' any package $2 and install it to $1. Replaces, can
# be set with $GO_GETTOOL_REPLACES.
function go.get-tool {
  local tmp_dir
  if ! [ -f "${1}" ]; then
    tmp_dir="$(mktemp -d)"
    # shellcheck disable=SC2064
    trap "rm -rf ${tmp_dir}" EXIT
    cd "$tmp_dir"
    go mod init tmp
    if [ -n "${GO_GETTOOL_REPLACES:-}" ]; then
      go mod edit -replace "${GO_GETTOOL_REPLACES}"
    fi
    echo "Downloading ${2}"
    GOBIN="$(dirname "${1}")" go get "${2}"
  fi
  echo "Using ${2} in ${1}"
}

# go.run-tool will 'go get' package $1 and install it to temp path. Replaces, 
# can be set with $GO_GETTOOL_REPLACES. You could pass arguments to the tool by
# setting $2 and up.
function go.run-tool {
  local path_id systmp_dir bin_dir package package_id bin
  package="${1:?'Pass a GO executable package ref as arg[1]'}"
  shift
  path_id="$(echo "${BASH_SOURCE[0]:-$0}" | sha1sum - | awk '{print $1}')"
  package_id="$(echo "$package" | sha1sum - | awk '{print $1}')"
  systmp_dir="$(dirname "$(mktemp -d -u)")"
  bin_dir="${systmp_dir}/go-run-tool/${path_id}/${package_id}"
  mkdir -p "$bin_dir"
  bin="${package%@*}"
  bin="${bin_dir}/${bin##*/}"

  go.get-tool "$bin" "$package"

  "$bin" "$@"
}