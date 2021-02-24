FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -mod=vendor -o api ./cmd/api

FROM alpine
WORKDIR /app
COPY --from=builder /app/api /bin/

EXPOSE 8080
CMD api