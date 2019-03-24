FROM golang as builder
WORKDIR /app
COPY * /app/
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
ENTRYPOINT ["/root/server"] 