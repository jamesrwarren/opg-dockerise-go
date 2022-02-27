## Technical Task

I have decided to go for a distroless static multistage build for my little go app.

The reason I have picked distroless static is because distroless gives 
added security (in the sense of less attack surface at least) over even alpine as 
it's completely bare bones and doesn't even have a shell (like alpine does). 

I used alpine as the builder because it's small compared to debian and other base images to speed up build times. 

As distroless has no shell at all, this could be a limiting factor when developing 
locally so I have provided an argument that you can override to use the debug version that has shell 
and tools from busybox.

I chose static over the other base images as it is described as ideal for statically compiled 
images like go and doesn't include the 
gcc libraries that base does as mentioned [here](https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md).

I chose to use built in user nonroot as this again makes the container more secure as any malicious code that 
is able to run in your app won't be running as root.

We use CGO_ENABLED=0 so that it creates statically linked binary meaning that all libraries are 
included in the binary and we don't dynamically link to external libraries once compiled. Apparently the `net` 
package we use in the example uses CGO so this ensures we're self-contained in our very barebones container.

As this is effectively a health check endpoint at the moment, I felt like it would be good to have something 
actually use it so I created the basic service in docker-compose and added a print 
(this could be real logging library in real app) and then added the healthcheck in the compose file. 
One issue though was that our image is so locked down that it didn't have any tool to use the healthcheck 
like curl or wget. wget is very small so we bring that in just for the healthcheck. 

## Getting started

You can test the app from outside the container by running:

```
go run main.go
```

and then:

```
curl -X GET localhost:8080/
```

To manually build an image and run it in debug mode (with a shell):

```
docker build -t opg-dockerise-go:latest --build-arg IMAGE_VERSION=debug .
```

or for production mode (which we would use in CI):

```
docker build -t opg-dockerise-go:latest .
```

Then to run the app we can do this:

```
docker run opg-dockerise-go:latest -p 8080:8080 -e APP_STATUS=OK
```

As this seems like the start of a bigger app and seems like it's a healthcheck currently so let's spin it up as a service 
and check it's getting healthcheck responses back (this also defaults to debug mode as this would be used in dev):

```
docker compose up
```

