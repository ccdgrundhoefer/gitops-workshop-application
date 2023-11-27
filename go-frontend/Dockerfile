# Stage 1: Build-Stage
FROM registry.access.redhat.com/ubi8/go-toolset:latest as builder

COPY cmd/ cmd/
COPY internal/ internal/

RUN go mod init frontend && \
    go mod tidy -e && \
    go build -o main ./cmd/frontend

# Stage 2: Production-Stage
FROM registry.access.redhat.com/ubi8/ubi-micro:latest

WORKDIR /app

# Kopiere die kompilierte Binärdatei aus der Build-Stage
COPY --from=builder /opt/app-root/src/main .
COPY static/ static/

# Exponiere den Port, auf dem die Anwendung läuft
EXPOSE 8080

# Starte die Anwendung
CMD ["./main"]
