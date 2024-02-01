FROM golang:1.21.6

WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /highcpu

ENTRYPOINT ["/highcpu"]
CMD ["-multi"]