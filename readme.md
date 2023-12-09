# Ozon Test Task

## Название микросервиса

Название микросервиса - test-microservice.
Для удобной проверки я сделал что-то типо фронта, который принимает http запросы
и общается с test-microservice через gRPC. Swagger доступен по `localhost:8080/swagger/index.html`.
Для запуска проекта можно использовать команду.
```
make up_build
```

## Выбор хранилища

Выбор хранилища зависит от переменных окружения. Для использования postgres необходимо добавить в окружение следующие переменные:

```
dbType: "postgres"
DSN: "host=postgres port=5432 user=postgres password=password dbname=urls sslmode=disable timezone=UTC connect_timeout=5"
dbURL: "postgres://postgres:password@postgres:5432/urls?sslmode=disable"
migrationPath: "file:///app"
```

Для запуска контейнера postgres с переменными среды используйте следующий пример [docker-compose.yml](./docker-compose.yml).

Если переменная dbType не равна postgres, то по умолчанию будет использоваться inMemory хранилище.

## Описание эндпоинтов

Микросервис реализует связь с другими микросервисами с помощью gRPC. Прото-файл находится в директории [proto](./test-microservice/protobuf/short/short.proto).

## Зависимости

Микросервис зависит от пакетов:
- [pgx](https://github.com/jackc/pgx/v4).
- [go-migrate](https://github.com/golang-migrate/migrate/v4)
- [gRPC](https://google.golang.org/grpc)
- [testify](https://github.com/stretchr/testify)
- [protobuf](https://google.golang.org/protobuf)
