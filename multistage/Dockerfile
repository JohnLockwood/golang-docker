FROM golang as builder
WORKDIR /app
ADD * /app/
RUN go build -o server . 

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
ENTRYPOINT ["/root/server"] 