FROM golang:1.20.9-bullseye AS base

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências necessárias
RUN go mod download

# Copie todo o código-fonte da aplicação para o diretório de trabalho
COPY . .

# Construa a aplicação
RUN go build -o main .

# Defina a porta que a aplicação vai escutar
EXPOSE 8447

# Comando para rodar o servidor
CMD ["./main"]


# Smaller build for prod
FROM debian:bullseye-slim AS release

