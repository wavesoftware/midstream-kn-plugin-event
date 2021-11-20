#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(realpath "$(dirname "${BASH_SOURCE[0]}")/../..")"
# shellcheck source=../lib.bash
source "${project_dir}/openshift/lib.bash"
tools="$(mktemp -d)"

go.get-tool "${tools}/deviate" \
  github.com/cardil/deviate/cmd/deviate@v0 \
  github.com/shurcooL/graphql=github.com/cli/shurcooL-graphql@v0.0.0-20200707151639-0f7232a2bf7e

"${tools}/deviate" sync --config "${project_dir}/.deviate.yaml"