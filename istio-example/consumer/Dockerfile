FROM golang:1.13

ENV GOPROXY=https://goproxy.io
ENV CGO_ENABLED=0

RUN mkdir /example-kube
WORKDIR /example-kube
ADD . /example-kube
RUN go build -o bin/consumer ./istio-example/consumer/main.go

FROM alpine:3.10

RUN mkdir /app
WORKDIR /app
COPY --from=0 /example-kube/bin/consumer .
ENTRYPOINT ["./consumer"]
EXPOSE 8081