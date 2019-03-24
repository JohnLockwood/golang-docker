# Golang-Docker

This is the source for the CodeSolid article, [Building a Golang Docker Container](http://codesolid.com/post/building-a-docker-golang-container/).

The article documents it better, but briefly...

We show two ways to build and run a go application, and document the approach in README files in the two folders:

* naive - This approach is simple but almost certainly not what you want for production.  It's worth starting here if you're somewhat new to Docker files, since we go over some basic commands and techniques.
* multistage - Here we use a multistage build to reduce the image size ***substantially*** -- from 700+ Mb down to 13, and remove the code from our final image.

