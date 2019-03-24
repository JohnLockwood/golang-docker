FROM golang
WORKDIR /app
ADD go.* /app/
ADD *.go /app/
RUN go build -o main . 
ENTRYPOINT ["/app/main"]
