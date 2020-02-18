#!/bin/bash

CURRENT_DIR="$(cd "$(dirname $0)"; pwd)"
EXAMPLE_HOME=${CURRENT_DIR}/../

docker build -f ${CURRENT_DIR}/consumer/Dockerfile -t consumer:1.0 ${EXAMPLE_HOME}
docker build -f ${CURRENT_DIR}/provider/Dockerfile -t provider:1.0 ${EXAMPLE_HOME}


