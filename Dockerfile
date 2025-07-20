FROM golang:1.23.10-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go env -w GOPROXY=https://proxy.golang.org,direct && \
    go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/bin/server ./server/cmd/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/bin/server /app/server

COPY db.sql /app/db.sql

EXPOSE 8080

CMD ["/app/server"]