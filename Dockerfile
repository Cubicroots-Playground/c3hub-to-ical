FROM golang:1.23.4-alpine3.20 as builder

WORKDIR /run

COPY ./ ./
RUN go mod download
RUN go build -o /run ./cmd/main.go

FROM alpine:3.21
COPY --from=builder /run/main /run/
WORKDIR /run

CMD ["/run/main"]