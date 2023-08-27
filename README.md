# Dynamic User Segmentation service

## API

## Структура управления сегменатцией

**POST - создать сегмент**

`/api/segment`

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

Ответ

````
{
	"deleted_segment": "AVITO_CAR_DISCOUNT",
	"message": "Completed successfully"
}
````