SRCS=webhook
ARCH ?= amd64
TARGET ?= bin
GOLANG_VERSION ?= latest
SRC_DIR = $(notdir $(shell pwd))
export

bin:
	../image-util.sh bin $(SRCS)

.PHONY: bin

image:
	docker build -t admission-webhook:1.0 ../

.PHONY: image

webhook: image
	kubectl apply -f webhook.yaml

.PHONY: webhook

mutating: webhook
	cat mutating-webhook-configuration.yaml | ../webhook-patch-ca-bundle.sh > mutating-webhook-configuration-ca-bundle-temp.yaml \
	&& kubectl apply -f mutating-webhook-configuration-ca-bundle-temp.yaml

.PHONY: mutating