package models

// структура таблицы машин
type Car struct {
	Id     string
	RegNum string
	Mark   string
	Model  string
	Year   string
	Owner  People
}
