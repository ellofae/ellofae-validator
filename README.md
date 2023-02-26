# ellofae-validator
___________________________________

## Валидатор, определяющий является ли переданная структура валидной, основываясь на переданных пользователем условиях валидности полей структуры.

Набор функций:

**func ValidateStruct(data interface{}, fields ...*FieldType) error** - определяет является ли переданная структура валидной, основываясь на переданных условиях валидности. Условия валидности входят в поле **rules** типа **FieldType**. Функция возвращает пользователю **определенную** ошибку, если структура не валидна, основываясь на переданных условиях валидности полей или вовзращает nil, если структура валидна.

Ошибка записывается в лог файл **logger.log** (путь: /tmp/logger.log)

**func ValidateStructInformative(data interface{}, fields ...*FieldType) error** - определяет является ли переданная структура валидной, основываясь на переданных условиях валидности, но служит в основном для записи всех возникших ошибок во время валидирования в лог файл. По завершению работы функции возвращает либо nil, если структура валидна, либо ошибку ErrStructNotValid.

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
