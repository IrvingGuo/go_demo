# Joynext Online App Http Server Demo

## Go

### CMD
1. run
   1. `go run cmd/main.go`
2. build and run binary(e.g. server)
   1. `go build -o ./out/server ./cmd/main.go`
   2. `./out/server`
   3. [Reference: Explanation of the above cmd](https://segmentfault.com/a/1190000013989448)
4. format code before commit `go fmt ./...`
5. set env vars
   1. [Reference Go Mirror speed up](https://learnku.com/go/wikis/38122): `go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct`

### 'go.mod' & 'go.sum'
[Reference: Understanding 'go.sum' and 'go.mod']((https://golangbyexample.com/go-mod-sum-module/))

## Docker

### Build with Dockerfile locally
[Multistage build](https://docs.docker.com/develop/develop-images/multistage-build/)
1. `docker build -t ronannnn/jrs-go:1.0 -f Dockerfile .`
   1. `-t`: tag
   2. `-f`: file
2. `docker tag ronannnn/jrs-go:1.0 ronannnn/jrs-go:latest`
3. `docker login` (login with username and password of docker hub)
4. `docker push ronannnn/jrs-go:1.0` & 

### Run locally
1. `docker run -d --rm -p 5001:5001 --name jrs 4242`
   1. `-d`: run in daemon
   2. `--rm`: delete container when stopping
   3. `-p 5001:5001`: map host 5001 to container 5001
   4. `4242`: first 4 numbers of image id
2. `docker logs fa11`
   1. `fa11`: container id
3. `docker stop/start fa11`

### Run locally with docker-compose
`docker-compose up/down/start/stop`, check the manual to see the difference

### Run Mysql with root password
`docker run --name mysql-jrs -e MYSQL_ROOT_PASSWORD=root -d -p 3306:3306 -v /var/mysql_data:/var/lib/mysql mysql`

### CMD vs Entrypoint
[Reference](https://stackoverflow.com/questions/21553353/what-is-the-difference-between-cmd-and-entrypoint-in-a-dockerfile)
1. The ENTRYPOINT specifies a command that will always be executed when the container starts.
2. The CMD specifies arguments that will be fed to the ENTRYPOINT
3. see examples in this link

### Build Go Image with caches of modules
[Reference](https://petomalina.medium.com/using-go-mod-download-to-speed-up-golang-docker-builds-707591336888)

### Specification of go build cmd in Dockerfile
`RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./out/server ./cmd/main.go`

[Reference: Explanation of the above cmd](https://segmentfault.com/a/1190000013989448)

[Another Reference of CGO_ENABLED](https://johng.cn/cgo-enabled-affect-go-static-compile/)

### Problems
1. Cannot find config.yaml
    1. If `WORKDIR` not specified, the default is `/`. 
    When executing `/app/server`, the config.yaml that is copied to `/app` cannot be found.
    2. Solution: Set `WORKDIR` or copy config to `/`

## GitHub/Gitlab Action
[Reference: Build Docker image and push](https://github.com/marketplace/actions/build-and-push-docker-images)
