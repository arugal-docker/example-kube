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

kube: image
	kubectl apply -f webhook.yaml

.PHONY: kube

validating:
	cat validatingwebhook.yaml | ../webhook-patch-ca-bundle.sh > validatingwebhook-ca-bundle-temp.yaml \
	&& kubectl apply -f validatingwebhook-ca-bundle-temp.yaml

.PHONY: validating