# Stage 1: Build-Stage
FROM docker.io/golang:1.21.4-alpine AS builder

WORKDIR /app

# Kopiere go.mod und go.sum in den Arbeitsverzeichnis
COPY go.mod .
COPY go.sum .

# Führe `go mod download` aus, um die Abhängigkeiten zu laden
RUN go mod download

# Kopiere den Rest des Codes in das Arbeitsverzeichnis
COPY . .

# Baue das Go-Anwendungsprogramm
RUN go build -o main ./cmd/frontend

# Stage 2: Production-Stage
FROM docker.io/alpine:latest

WORKDIR /app

# Kopiere die kompilierte Binärdatei aus der Build-Stage
COPY --from=builder /app/main .

# Exponiere den Port, auf dem die Anwendung läuft
EXPOSE 8080

# Starte die Anwendung
CMD ["./main"]
