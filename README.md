# ellofae-validator
___________________________________

## Валидатор, определяющий является ли переданная структура валидной, основываясь на переданных пользователем условиях валидности полей структуры.

Набор функций:

**func ValidateStruct(data interface{}, fields ...*FieldType) error** - определяет является ли переданная структура валидной, основываясь на переданных условиях валидности. Условия валидности входят в поле **rules** типа **FieldType**. Функция возвращает пользователю **определенную** ошибку, если структура не валидна, основываясь на переданных условиях валидности полей или вовзращает nil, если структура валидна.

Ошибка записывается в лог файл **logger.log** (локальный путь: ./tmp/logger.log)

**func ValidateStructInformative(data interface{}, fields ...*FieldType) error** - определяет является ли переданная структура валидной, основываясь на переданных условиях валидности, но служит в основном для записи всех возникших ошибок во время валидирования в лог файл. По завершению работы функция возвращает либо nil, если структура валидна, либо ошибку ErrStructNotValid.

## Все возможные ошибки во время валидации:
  	ErrNilStruct                        структура данных должна быть определена, а не равняться nil
  
	ErrLogFileNotOpened                 лог файл не был успешно открыт
  
	ErrStructNotValid                   структура не валидна
  
	ErrValueNotUnsignedInt              возможные типы переданных значений - uint8, uint16, uint32, uint64, uint
  
	ErrRegexNotSatisfied                regex сравение не прошло успешно
  
	ErrNotValidatable                   типом данных должен быть указатель, структура или nil
  
	ErrFieldNotSpecified                зачение поля не может быть nil из-за требования к значению поля быть non-nil
  
	ErrOnlyStringValue                  значение должно быть строкой
  
	ErrStringLengthIsNotSatisfied       длина строки вызодит за пределы минимального или максимального значения
	
## Условия валидации полей, доступные пользователю:
Все доступные пользователю условия валидации полей структур реализуют метод **Specifier(value interface{}) error**. Так как условия валидации реализуют данный метод, то они могут быть представлены в виде типа **Rule**, служащим основным типом(интерфейсом) условий, во время передачи их в функцию **Field(field interface{}, rules ...Rule) *FieldType**. При определения собственных условий валидации, тип должен реализовывать метод **Specifier(value interface{}) error** 

Доступные пользователю условия валидации:
	Required    			зачение поля должно быть определено (не быть nil)
	
	UnsignedInt 			тип значения поля uint8, uint16, uint32, uint64, uint
	
	Length struct			длина строки должна быть в интервале между минимальным и максимальным значениями
		MinValue int
		MaxValue int
	}
	
	MatchRequired struct {		строка должна удовлетворять заданному пользователем regex условию
		RegexToMatch string
	}
	

## Валидация структур

Структура должна быть передана функцям в виде указателя на нее, иначе будет возвращена ошибка ErrNotValidatable
Пример использования функции **func ValidateStruct(data interface{}, fields ...*FieldType) error**:

```
type MyType struct {
	IntField    int64
	StringField string
	NilPtrField *int
}

myStruct := MyType{255, "test string", nil}

err := validation.ValidateStruct(&myStruct,
		validation.Field(&myStruct.IntField, validation.UnsignedInt),
		validation.Field(&myStruct.StringField, validation.MatchRequired{"[A-Z]"}, validation.Length{4, 20}),
		validation.Field(&myStruct.NilPtrField, validation.Required))

if err != nil {
	// Обработка любой возникшей ошибки
	fmt.Println("Result: Not valid... check the log file '../tmp/logger.log' for more information")
}
```

Пример использования функции **func ValidateStructInformative(data interface{}, fields ...*FieldType) error**:

```
err = validation.ValidateStructInformative(&myStruct,
		validation.Field(&myStruct.IntField, validation.UnsignedInt),
		validation.Field(&myStruct.StringField, validation.MatchRequired{"[A-Z]"}, validation.Length{10, 20}),
		validation.Field(&myStruct.NilPtrField, validation.Required))

if err != nil {
	fmt.Println(err) // обработка ошибки ErrStructNotValid
}
```
