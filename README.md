# Сервис динамического сегментирования пользователей

Сервис написан на Go с использованием фреймфорка [fiber](https://github.com/gofiber/fiber).\
В docker compose развернута PostgreSQL 15 вместе с pgAdmin.\
Для взаимодействия с БД используется драйвер [pgx](https://github.com/jackc/pgx).\
Для доступа к переменным окружения из .env файла используется библиотека [godotenv](https://github.com/joho/godotenv).


## API 
Взаимодействие с сервисом осуществляется по REST API.\
В репозитории есть файл SegmentsServiceAPI.yaml с описанием API (по спецификации OpenAPI 3.0).

[Также можно посмотреть API на SwaggerHub](https://app.swaggerhub.com/apis/echpochmak31/AvitoTechSwagger/1.0.0)

## Примеры запросов и ответов (через Postman)

