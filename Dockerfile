# Build in go env
FROM golang:1.17.2-alpine3.14 AS builder
WORKDIR /go/src/
# Go mod will be cached in this way (if mod/sum is not modified)
COPY go.mod .
COPY go.sum .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
# Copy remaining source files
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./launchserver ./cmd/main.go

# Start fresh from a smaller image
# debug in 'alpine' if needed and deploy in 'scratch'
FROM scratch
# FROM alpine
WORKDIR /app/
# if certificates needed
# RUN apk add --update --no-cache ca-certificates git
COPY --from=builder /go/src/launchserver ./
COPY config.yaml .
# COPY ["Master RD Resource Plan.xlsx", "."]
# RUN ["chmod", "+x", "./launchserver"]
CMD ["./launchserver"]
