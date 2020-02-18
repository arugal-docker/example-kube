#!/bin/bash

CURRENT_DIR="$(cd "$(dirname $0)"; pwd)"
EXAMPLE_HOME=${CURRENT_DIR}/../

kubectl delete -f ${CURRENT_DIR}/istio-example.yaml
kubectl delete -f ${CURRENT_DIR}/istio-example-ingress.yaml

docker rmi consumer:1.0
docker rmi provider:1.0

docker system prune