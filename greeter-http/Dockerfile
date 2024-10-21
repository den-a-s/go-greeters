FROM golang:1.22.7

ENV SERVICE_NAME=greeter-http

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/${SERVICE_NAME} ./cmd/${SERVICE_NAME}

EXPOSE 8080
# Интерполяция не работает в exec режиме (массив аргументов)
CMD ./bin/$SERVICE_NAME