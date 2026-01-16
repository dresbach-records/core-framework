# Usa a imagem oficial do Go com Alpine Linux (uma versão leve)
FROM golang:1.22-alpine

# Instala o 'air', uma ferramenta para live-reloading de aplicações Go.
# Isso permite que nosso servidor reinicie automaticamente quando o código for alterado.
RUN go install github.com/cosmtrek/air@latest

# Define o diretório de trabalho dentro do contêiner.
WORKDIR /app

# Copia os arquivos de gerenciamento de dependências.
COPY go.mod go.sum ./

# Baixa as dependências do Go para otimizar o cache de layers do Docker.
RUN go mod download

# Copia todo o resto do código do nosso projeto para o contêiner.
COPY . .

# Expõe a porta 8080 para que possamos acessá-la de fora do contêiner.
EXPOSE 8080

# O comando que será executado quando o contêiner iniciar.
# Ele inicia o 'air' que, por sua vez, compila e executa nossa aplicação.
CMD ["air", "-c", ".air.toml"]