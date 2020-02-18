FROM golang:1.13

ENV GOPROXY=https://goproxy.io
ENV CGO_ENABLED=0

RUN mkdir /example-kube
WORKDIR /example-kube
ADD . /example-kube
RUN go build -o bin/provider ./istio-example/provider/main.go

FROM alpine:3.10

RUN mkdir /app
WORKDIR /app
COPY --from=0 /example-kube/bin/provider .
ENTRYPOINT ["./provider"]
EXPOSE 8080