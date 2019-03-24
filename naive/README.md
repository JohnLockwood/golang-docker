# Docker Go Hello-World Web Server Example

This is a simple example of how to build and run a Go application in Docker.  In our case it's a simple web server running "hello world".

This is a one step build and run, which could be improved to exclude go from the final container to make the resulting image smaller.

## Building

Build the image tagging it with the name gorilla-demo using the following command:

```
docker build -t golang-gorilla:latest .
```

## Running

Run the image in detached mode, naming the container "gorilla", and mapping the container port (9000) to a host port (8080):

```
docker run --name gorilla -p8080:9000 -d golang-gorilla
```

With this running you can then reach the server from the host:

```
curl localhost:8080
```

Or from the container:

```
docker exec -it gorilla curl http://localhost:9000
```