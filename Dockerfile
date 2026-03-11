# Etapa 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# Compila o binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -o app-a .

# Etapa 2: Imagem final (leve)
FROM alpine:latest
WORKDIR /root/
# Copia apenas o binário da etapa anterior
COPY --from=builder /app/app-a .
EXPOSE 8080
CMD ["./app-a"]