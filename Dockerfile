# Используем официальный образ Golang как базовый
FROM golang:latest AS builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем исходный код в контейнер
COPY . .

# Собираем и компилируем наше приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api/main.go

# Используем scratch для минимального образа
FROM scratch
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
