SRCS=webhook
ARCH ?= amd64
TARGET ?= bin
GOLANG_VERSION ?= latest
SRC_DIR = $(notdir $(shell pwd))
export


bin:
	../image-util.sh bin $(SRCS)

.PHONY: bin