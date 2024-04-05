package models

// структура таблицы владельцев
type People struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
