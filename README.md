# Сервис динамического сегментирования пользователей


## Описание:

Сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

## Используемые технологии:

- PostgreSQL (в качестве хранилища данных)
- Docker (для запуска сервиса)
- Библиотека sqlx для работы с БД
- Регистрация и аутентификация для создания юзеров. Работа с JWT.
- Фреймворк gin-gonic/gin

## Запуск

Запустить сервис можно с помощью команды `docker-compose up --build avito_test`
Для миграции использовать следующую команду `migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up`


## Использование

Примеры того, как использовать ваш проект.

## Регистрация пользователя

### Пример запроса:

`POST http://localhost:8000/auth/sign-up`
`Content-Type: application/json`
```
{
    "name":"Uliana",
    "username":"ulnest",
    "password":"qwerty"
}
```

### Пример успешного ответа:
```
"id": 2
```

## Авторизация пользователя

### Пример запроса:

`POST http://localhost:8000/auth/sign-in`
`Content-Type: application/json`
```
{
    "username":"ulnest",
    "password":"qwerty"
}
```

### Пример успешного ответа:
```
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM1NTM1MDksImlhdCI6MTY5MzUxMDMwOSwidXNlcl9pZCI6MX0.amFa_PHMY8QA7dCeZY8IPooZ6KQxPltnKHrJu9QyP_c"
```

## Создание сегмента

### Пример запроса:

`POST http://localhost:8000/api/segments`
`Content-Type: application/json`
```
{
    "slug":"AVITO_DISCOUNT_50"
}
```

### Пример успешного ответа:
```
 "id": 1
```

## Удаление сегмента

### Пример запроса:

`DELETE http://localhost:8000/api/segments/AVITO_DISCOUNT_50`
`Content-Type: application/json`


### Пример успешного ответа:
```
 "status": "ok"
```

## Добавление пользователя в сегмент

### Пример запроса:

`POST http://localhost:8000/api/users/2/segments`
`Content-Type: application/json`
```
{
    "slug":"AVITO_DISCOUNT_50"
}
```

### Пример успешного ответа:
```
 "status": "ok"
```

## Удаление пользователя из сегмента

### Пример запроса:

`http://localhost:8000/api/users/2/AVITO_DISCOUNT_50`
`Content-Type: application/json`

### Пример успешного ответа:
```
 "status": "ok"
```

## Получение сегментов, в которых состоит пользователь

### Пример запроса:

`GET http://localhost:8000/api/users/2/segments`
`Content-Type: application/json`


### Пример успешного ответа:
```
 "segments": [
        "AVITO_DISCOUNT_50"
    ]
```