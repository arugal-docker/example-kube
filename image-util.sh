#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

TASK=$1


# This function is for building the go code
bin() {
  for SRC in $@;
  do
    go build -a -o ${TARGET}/${SRC} ./$(dirname ${SRC})
  done
}

shift

eval ${TASK} "$@"