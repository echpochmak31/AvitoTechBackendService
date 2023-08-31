# Сервис динамического сегментирования пользователей

Сервис написан на Go с использованием фреймфорка [fiber](https://github.com/gofiber/fiber).\
В docker compose развернута PostgreSQL 15 вместе с pgAdmin.\
Для взаимодействия с БД используется драйвер [pgx](https://github.com/jackc/pgx).

## Установка
1. Убедитесь, что установлена 1.18 версия Go или выше
2. Создайте свой проект и используйте ``go mod init github.com/your/repo``
3. Запустите команду ``go get github.com/echpochmak31/avitotechbackendservice@main``
4. Поднимите docker compose

## Getting started
Прежде убедитесь, что все [необходимые переменные окружения](#переменные-окружения) доступны.

Для запуска серсиа достаточно вызвать функцию RunSegmentsService из пакета application 
и передать в нее адрес, по которому будут посылаться запросы.
```go
...
err := application.RunSegmentsService("127.0.0.1:8080")
...
```
Пример с использованием библиотеки [godotenv](https://github.com/joho/godotenv) для получения переменных окружения из .env файла. 
```
package main

import (
	"github.com/echpochmak31/avitotechbackendservice/pkg/application"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	err := application.RunSegmentsService("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}

```

## API 
Взаимодействие с сервисом осуществляется через REST API.\
В репозитории есть файл SegmentsServiceAPI.yaml с описанием API (по спецификации OpenAPI 3.0).

[Также можно посмотреть API на SwaggerHub](https://app.swaggerhub.com/apis/echpochmak31/AvitoTechSwagger/1.0.0)

#### Пример get запроса на получение сегментов пользователя с id 4:
Запрос через curl:\
``curl localhost:8080/segments/user/4``\
Ответ:\
``{"SegmentSlugs":["AVITO_PERFORMANCE_VAS","AVITO_VOICE_MESSAGES","AVITO_DISCOUNT_30"]}``

## Переменные окружения
1. Логин и пароль для PostgreSQL
```
POSTGRES_USER=user
POSTGRES_PASSWORD=pass
```
2. Маппинг портов для PostgreSQL
```
# Порт внутри docker
POSTGRES_PORT_SRC=5432

# Открытый порт, по которому будет происходить соединение
POSTGRES_PORT_DST=5433 
```
3. Маунт директорий для PostgreSQL И pgAdmin
```
POSTGRES_VOLUME=C:/path/to/postgresql_dir
PGADMIN_VOLUME=C:/path/to/pgadmin_dir 
```
4. Логин и пароль для pgAdmin
```
PGADMIN_USER=admin@admin.com
PGADMIN_PASSWORD=secret
```
5. Хост и имя БД для PostgreSQL
```
DB_HOST=localhost
DB_NAME=segments_service
```
6. Маунт директорий для отчетов
```
# Реальный путь, куда будут сохраняться отчеты
PATH_TO_REPORTS=C:/path/to/reports

# Путь в контейнере PostgreSQL
VIRTUAL_PATH_TO_REPORTS=/tmp/reports
```