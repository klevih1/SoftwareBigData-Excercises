# === Stage 1: Builder ===
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN chmod +x /app/scripts/build-application.sh
RUN /app/scripts/build-application.sh

# === Stage 2: Runner ===
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ordersystem .

EXPOSE 3000

CMD ["/app/ordersystem"]

