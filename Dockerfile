# MULTI STAGE EXAMPLE
FROM golang:1.24 AS build

ENV CGO_ENABLED=0
ENV GOOS=linux

# Устанавливаем место назначения для COPY
WORKDIR /app

# Копируем файлы go.mod и go.sum в WORKDIR
COPY go.mod go.sum ./
# Скачиваем необходимые Go модули (зависимости нашего проетка)
RUN go mod download

# Копируем все исходные go файлы нашего проекта в образ
# https://docs.docker.com/reference/dockerfile/#copy
ARG SERVICE
COPY ./cmd/${SERVICE}/*.go ./
# Собираем бинарный файл нашего приложения
RUN go build -o /bin/main 

#######################################
# STAGE 2. FINAL STAGE
#######################################

FROM alpine:3.19 AS final

RUN apk add --no-cache curl

WORKDIR /

COPY --from=build /bin/main /main

ENV PORT=8080
EXPOSE 8080

# Точка входа
ENTRYPOINT ["/main"]