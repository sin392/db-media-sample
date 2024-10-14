FROM golang:latest AS base
WORKDIR /app
COPY . .

FROM base AS debugger
WORKDIR /app
COPY . .
RUN go get github.com/go-delve/delve/cmd/dlv
EXPOSE 3001 40000
ENTRYPOINT ["dlv", "debug", "--listen=:40000", "--headless", "--accept-multiclient", "--continue", "--api-version=2"]

FROM base AS development
WORKDIR /app
COPY . .
RUN go install github.com/movio/bramble/cmd/bramble@latest
EXPOSE 8082
ENTRYPOINT ["bramble"]
CMD ["--config", "config.json"]

# ここいらないかも？
FROM base AS builder
WORKDIR /app
COPY . .
RUN go build -a --installsuffix cgo --ldflags="-s" -o main

FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
RUN go install github.com/movio/bramble/cmd/bramble@latest
COPY --from=builder /app .
EXPOSE 8082
ENTRYPOINT ["bramble"]
CMD ["--config", "config.json"]
