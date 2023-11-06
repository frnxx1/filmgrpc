# Используем образ golang в качестве базового образа
FROM golang:1.21.1-alpine AS build

# Устанавливаем зависимости
WORKDIR /filmgrpc
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем клиент
WORKDIR /filmgrpc/client
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client .

# Собираем сервер
WORKDIR /filmgrpc/service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Используем образ scratch в качестве базового образа для финального образа
FROM scratch

# Копируем бинарники из предыдущего этапа
COPY --from=build /filmgrpc/service/server /filmgrpc/service/server
COPY --from=build /filmgrpc/client/client /filmgrpc/client/client

# Запускаем сервер по умолчанию при запуске контейнера
CMD ["/filmgrpc/service/server"]