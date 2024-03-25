# Используем образ golang в качестве базового образа
FROM golang:1.17-alpine AS build

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем зависимости Go модуля
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код внутрь контейнера
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/api

# Отдельный этап сборки
FROM alpine:latest

# Устанавливаем необходимые пакеты
RUN apk --no-cache add ca-certificates

# Копируем бинарный файл из предыдущего этапа сборки
COPY --from=build /app/main /usr/local/bin/main

# Устанавливаем переменную окружения для порта
ENV PORT=8080

# Определяем порт, который приложение будет использовать
EXPOSE $PORT

# Запускаем приложение
CMD ["main"]
