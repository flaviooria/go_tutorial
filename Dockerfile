FROM go:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd/cli/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]