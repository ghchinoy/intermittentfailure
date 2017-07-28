# Development

Build for Dockerfile (ref @kelseyhightower's [Optimizing docker images for go](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07))

```
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
docker build -t ghchinoy/intermittentfailure .
```

Run, exposing on random port (`-P`)

```
docker run -d -P ghchinoy/intermittentfailure
```

on gce, using `slowendpoint` with container-optimized image

```
docker pull ghchinoy/intermittentfailure
docker run -d -p 80:8080 --name failureapi ghchinoy/intermittentfailure
```