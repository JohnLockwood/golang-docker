# Docker Go Hello-World Web Server Example

Viewing the image size from the naive build:

```
docker images golang-gorilla
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
golang-gorilla      latest              67edf2eee9fa        8 hours ago         782MB
```

To optimize it, we can use a [multi-stage build](https://docs.docker.com/develop/develop-images/multistage-build/), a Docker feature that allows us to create a build using a preliminary image, then package the results into a much smaller -- and much more secure -- final image.

## Building

We build it as we built our earlier image, but since we expect it to be smaller, we'll give it a new name, mini-gorilla.  

```
docker build -t mini-gorilla:latest .
```

Let's see if we saved any space:

```
docker images mini-gorilla
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
mini-gorilla        latest              5ae087964c8d        53 seconds ago      13.3MB
```


## Running

Run the image in detached mode, naming the container "gorilla", and mapping the container port (9000) to a host port (8080):

```
docker run --name gorilla2 -p8080:9000 -d mini-gorilla
```

With this running you can then reach the server from the host:

```
curl localhost:8080
```

Or from the container:

```
docker exec -it gorilla curl http://localhost:9000
```