FROM golang:1.22.0 AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM scratch

COPY --from=builder /app/main /

EXPOSE 8080

ENTRYPOINT ["/main"]
