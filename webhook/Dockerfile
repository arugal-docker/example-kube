FROM golang:1.13

ENV GOPROXY=https://goproxy.io
ENV CGO_ENABLED=0

RUN mkdir /example-kube

ADD . /example-kube/

WORKDIR /example-kube/webhook
RUN make bin

FROM alpine:3.10
COPY  --from=0 /example-kube/webhook/bin/webhook /webhook
ENTRYPOINT ["/webhook"]