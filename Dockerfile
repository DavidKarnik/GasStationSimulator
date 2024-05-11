# Použití oficiálního Golang image z Docker Hub
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="David Karnik <david.karnik@tul.cz>"

# Nastavení pracovního adresáře v kontejneru
WORKDIR /app

# Copy go mod and sum files <- must have .mod
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# Zkopírování obsahu Go projektu do kontejneru
COPY . .

# Sestavení Go programu
RUN go build -o main .

# Určení, který port bude expozován
EXPOSE 8080

# Příkaz, který se spustí při spuštění kontejneru
CMD ["./main"]
