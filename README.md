# Dynamic User Segmentation service

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

Выполнено:

- [x] Метод создания сегмента;
- [x] Метод удаления сегмента;
- [x] Метод получения списка всех сегментов;
- [x] Метод добавления пользователя в сегмент;
- [x] Метод удаления сегментов у пользователя;
- [x] Метод получения активных сегментов пользователя;
- [ ] Метод получения списка всех пользователей с их сигментами;
- [x] Swagger;
- [ ] Тестирование кода;
- [ ] Функционал с CSV;
- [ ] TTL сегментам;
- [ ] % выпадания пользователям сегмента;