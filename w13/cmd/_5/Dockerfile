# Этап сборки
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Копирование и загрузка зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование и сборка кода
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api

# Этап создания минимального образа
FROM alpine:latest

# Установка необходимых пакетов
RUN apk --no-cache add ca-certificates tzdata

# Установка временной зоны
ENV TZ=Asia/Almaty

WORKDIR /root/

# Копирование бинарного файла из этапа сборки
COPY --from=builder /app/api .
COPY --from=builder /app/configs ./configs

# Экспорт порта
EXPOSE 8080

# Запуск приложения
CMD ["./api"]