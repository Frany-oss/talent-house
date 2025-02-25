# Reto Tecnico Talent House

El proyecto se desplego en DockerHub para que sea rapido y facil poder testear e instalar las apis.

## Api en GO 
Correr en terminal
```bash
docker run -d -p 8080:8080 --name go-api -e NODE_API_URL=http://node-api:3000 franyy/talent-house-go-api:latest
```

## Api en Express
Correr en terminal
```bash
docker run -d -p 3000:3000 --name node-api -e NODE_API_URL=http://node-api:3000 franyy/talent-house-node-api:latest
```

## Como probar el funcionamiento
### 1. Endpoint para la autenticación con JWT
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "username": "admin",
    "password": "password"
}'
```
Ese endpoint generará un token el cual es necesario tanto para hacer la factorización QR como para obtener las operaciones adicionales en el otro api

### 2. Endpoint para enviar la matriz - da como resultado los valores ya procesados por el api de Express

```bash
curl -X POST http://localhost:8080/qr \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <jwt-token>" \
-d '{
    "matrix": [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
}'
```

## Unit Test
### Api en Go
```bash
go test ./...
```
### Api en Express
```bash
npm test
```
