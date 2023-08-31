# Dynamic User Segmentation service

## Инструкции по запуску:

`make server` - Запуск файла main.go

`make dev` - Поднятие dev-среды



## API

## Структура управления сегменатцией

**POST - создать сегмент**

`/api/segment`

Запрос

````
{
	"segment":"avito_car_discount"
}
````

Ответ

````
{
	"created_segment": "AVITO_CAR_DISCOUNT",
	"message": "Completed successfully"
}
````

**GET - получить все сегменты**

`/api/segment/`

Ответ

````
{
    "1": "AVITO_VOICE_MESSAGES",
    "2": "AVITO_PERFORMANCE_VAS",
    "3": "AVITO_DISCOUNT_30",
    "4": "AVITO_DISCOUNT_50"
}
````

**DELETE - удалить сегмент**

`/api/segment/:segment`

Запрос

````
{
	"segment":"AVITO_CAR_DISCOUNT"
}
````

Ответ

````
{
	"deleted_segment": "AVITO_CAR_DISCOUNT",
	"message": "Completed successfully"
}
````

## Структура управления сегменатцией у пользователей

**POST - получить все сегменты пользователя**

`/api/user/:id`

Запрос 

```
{
	"id": "0fa0c12b-32a5-49ef-bbb5-3539f8f67971"
}
```

Ответ 

```
{
	"id": "0fa0c12b-32a5-49ef-bbb5-3539f8f67971",
	"segments": [
		{
			"segment": "AVITO_VOICE_MESSAGES"
		},
		{
			"segment": "AVITO_DISCOUNT_50"
		}
	]
}
```

**POST - добавить сегменты пользователю**

`/api/user/`

Запрос 

```
{
	"id": "0fa0c12b-32a5-49ef-bbb5-3539f8f67971",
	"segments": [
		{
			"segment": "AVITO_VOICE_MESSAGES"
		},
		{
			"segment": "AVITO_DISCOUNT_50"
		}
	]
}
```

Ответ 

```
{
	"added_segments": [
		{
			"id": 1,
			"segment": "AVITO_VOICE_MESSAGES"
		},
		{
			"id": 3,
			"segment": "AVITO_DISCOUNT_50"
		}
	],
	"message": "Completed successfully"
}
```

**DELETE - удалить сегменты у пользователю**

`/api/user/:segment`

Запрос

```
{
	"id": "70c247da-377a-42ac-97f6-316abfc43722",
	"segments": [
		{
			"segment": "AVITO_VOICE_MESSAGES"
		},
		{
			"segment": "AVITO_DISCOUNT_50"
		}
	]
}
```

Ответ 

```
{
	"deleted_segments": [
		{
			"id": 1,
			"segment": "AVITO_VOICE_MESSAGES"
		},
		{
			"id": 3,
			"segment": "AVITO_DISCOUNT_50"
		}
	],
	"message": "Completed successfully"
}
```

**GET - получение всех пользователей**

`/api/user/all`

Ответ

```
{
	"0fa0c12b-32a5-49ef-bbb5-3539f8f67971": [
		"AVITO_carrR_FFGD"
	],
	"59421123-416e-451f-96ca-c8a5475ff210": [
		"AVITO_VOICE_MESSAGES",
		"AVITO_CREATE_SEGMENT"
	],
	"974b7ce1-c719-45f4-933d-e73376bdb990": [
		"null"
	],
	"9ea5d9c7-fa09-42d1-8534-21d7e0eac8cc": [
		"AVITO_DICS_TOP",
		"AVITO_VOICE_MESSAGES",
		"AVITO_carrR_FFGD",
		"AVITO_CAJARRR_FFGD"
	]
}
```

Выполнено:

- [x] Метод создания сегмента;
- [x] Метод удаления сегмента;
- [x] Метод получения списка всех сегментов;
- [x] Метод добавления пользователя в сегмент;
- [x] Метод удаления сегментов у пользователя;
- [x] Метод получения активных сегментов пользователя;
- [x] Метод получения списка всех пользователей с их сигментами;
- [x] Swagger;
- [ ] Тестирование кода;
- [ ] Функционал с CSV;
- [ ] TTL сегментам;
- [ ] % выпадания пользователям сегмента.


## Уточнение

Задание связанное с CSV файлом выполнено на половину. Реализация данного задания следующая:
при каждом взаимодействии с пользователем (добавление сегмента, удаление сегмента),
он попадает в общий файл  `usersLogs.csv`.
Также было в планах реализовать метод,
через который можно будет получить ссылку на следующий файл `userID_segment.csv`. А именно через яндекс API 
и  облачное хранилище. 

На успешное выполнение дополнительных заданий не хватило времени.

