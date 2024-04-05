package models

// структура таблицы машин
type Car struct {
	Id     string `json:"id"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   string `json:"year"`
	Owner  People `json:"owner"`
}
