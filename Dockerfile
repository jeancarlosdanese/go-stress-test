# Etapa de build
FROM golang:1.22-alpine AS builder

# Define diretório de trabalho
WORKDIR /app

# Copia os arquivos de dependência e baixa os módulos
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código da aplicação
COPY . .

# Compila o binário estático
RUN go build -o stress-tester main.go

# Etapa final (imagem pequena para execução)
FROM alpine:latest

WORKDIR /app

# Copia apenas o binário da etapa de build
COPY --from=builder /app/stress-tester .

# Define o entrypoint
ENTRYPOINT ["/app/stress-tester"]
