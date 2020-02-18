SRCS=istio-example
ARCH ?= amd64
TARGET ?= bin
GOLANG_VERSION ?= latest
SRC_DIR = $(notdir $(shell pwd))
export


image:
	./setup.sh

.PHONY: image

istio-example:
	kubectl apply -f istio-example.yaml && \
		kubectl apply -f istio-example-ingress.yaml

.PHONY: istio-example

clean:
	kubectl delete -f istio-example.yaml && \
		kubectl delete -f istio-example-ingress.yaml

.PHONY: clean
