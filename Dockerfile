# Stage 1: Build the Go application
FROM golang:latest AS builder

WORKDIR /app

# Копируем только go модули для ускорения сборки
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все остальные файлы
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/app/main.go

# Stage 2: Create the final image
FROM alpine:latest

# Устанавливаем необходимые зависимости
RUN apk --no-cache add postgresql redis

# Копируем скомпилированный бинарный файл из первого этапа
COPY --from=builder /app/main .

# Открываем порты для postgres и redis
EXPOSE 5432 6379

# Запускаем приложение
CMD ["./main"]
