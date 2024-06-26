# Используем официальный образ Golang как базовый
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod ./
COPY go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код проекта в рабочую директорию
COPY . .

# Собираем приложение для продакшена
RUN go build -o main .

# Определяем порт, на котором будет работать приложение
EXPOSE 8080

# Запускаем скомпилированный файл при запуске контейнера
CMD ["./main"]
