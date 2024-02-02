FROM golang:1.21.6-bookworm
COPY . /go/src/github.com/alphauslabs/hellok8s/
WORKDIR /go/src/github.com/alphauslabs/hellok8s/
RUN CGO_ENABLED=0 GOOS=linux go build -v -trimpath -installsuffix cgo -o hellok8s .

FROM debian:stable-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /hellok8s/
COPY --from=0 /go/src/github.com/alphauslabs/hellok8s/hellok8s .
# ENTRYPOINT ["/hellok8s/hellok8s"]
CMD ["/hellok8s/hellok8s"]
