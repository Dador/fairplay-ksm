FROM golang:1.17.13-buster AS builder
LABEL org.opencontainers.image.source https://github.com/Dador/fairplay-ksm

WORKDIR /app
COPY . /app

# Test
WORKDIR /app/ksm
RUN go get -u -t ./... && go test -run ''

WORKDIR /app/cryptos
RUN go get -u -t ./... && go test -run ''

# Build
WORKDIR /app
# ENTRYPOINT ["tail", "-f", "/dev/null"]
RUN go build -o /app/ksm-server ./api && chmod 755 /app/ksm-server

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /app/ksm-server .

EXPOSE 8080
ENTRYPOINT [ "/app/ksm-server" ]
