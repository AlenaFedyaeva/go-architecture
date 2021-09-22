package repository

//ItemFilter -filter items in DB
type ItemFilter struct {
	PriceLeft  *int64 //Опциональные парамеры делаем указателями
	PriceRight *int64
	Limit      int //На сколько сместились от 0 элемента
	Offset     int //Сколько элементов на странице
}
