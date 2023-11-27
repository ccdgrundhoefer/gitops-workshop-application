# Stage 1: Build-Stage
FROM registry.access.redhat.com/ubi8/go-toolset:latest as builder

WORKDIR /app

# Kopiere den Rest des Codes in das Arbeitsverzeichnis
COPY cmd/ cmd/
COPY internal/ internal/
COPY static/ static/

# Baue das Go-Anwendungsprogramm
RUN go mod init go-frontend && \
    go mod tidy && \
    go build -o main ./cmd/frontend

# Stage 2: Production-Stage
FROM registry.access.redhat.com/ubi8/ubi-micro:latest

WORKDIR /app

# Kopiere die kompilierte Binärdatei aus der Build-Stage
COPY --from=builder /app/main .

# Exponiere den Port, auf dem die Anwendung läuft
EXPOSE 8080

# Starte die Anwendung
CMD ["./main"]
