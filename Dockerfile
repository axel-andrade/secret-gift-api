# Use uma imagem mínima do Go como imagem base, como a imagem Alpine
FROM golang:1.24.0-alpine AS builder

# Defina variáveis de ambiente para otimizar a compilação
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Crie um diretório de trabalho dentro do contêiner
WORKDIR /app

# Instale o GCC e outras ferramentas de compilação necessárias
RUN apk --no-cache add build-base

# Copie apenas o arquivo go.mod e go.sum para permitir o cache das dependências
COPY go.mod go.sum ./

# Baixe e instale as dependências, incluindo atualizar o go.sum e baixar dependências ausentes
RUN go mod tidy

# Copie o restante dos arquivos de origem para o diretório de trabalho
COPY . .

# Compile a aplicação Go
RUN go build -o app ./cmd/api/main.go

# Use uma imagem base mais leve para a execução da aplicação
FROM alpine:latest

# Instale apenas os certificados CA necessários para comunicação segura (se necessário)
RUN apk --no-cache add ca-certificates

# Copie o binário compilado da etapa anterior
COPY --from=builder /app/app /app/

# Exponha a porta em que a aplicação irá rodar
EXPOSE 8080

# Comando para executar a aplicação
CMD ["/app/app"]
