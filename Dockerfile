# Start from Golang v1.14 base image to build the server
FROM golang:1.14.4-alpine3.12 as build

# ENV GO111MODULE=on

# Install git & Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Change workdir
WORKDIR /app

# Clone this repo
RUN git clone https://github.com/OhMinsSup/connecting-server.git /app

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 의존성 mod / sum을 변경하지 않으면 캐시됩니다.
RUN go mod download

# Build server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o connecting-server  .

# Run the server in this container #
FROM alpine:3.12

WORKDIR /app

COPY --from=build /app/.env.dev .
COPY --from=build /app/connecting-server .

CMD ["./connecting-server"]
