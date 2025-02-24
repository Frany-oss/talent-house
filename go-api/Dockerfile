# Utiliza la imagen oficial de Golang como imagen base
FROM golang:1.23

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de módulos de Go
COPY go.mod go.sum ./

# Descarga las dependencias especificadas en los archivos de módulos
RUN go mod download

# Copia el código fuente de la aplicación al contenedor
COPY . .

# Compila la aplicación Go y genera un ejecutable llamado "main"
RUN go build -o main .

# Expone el puerto en el que se ejecuta la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación cuando el contenedor se inicie
CMD ["./main"]