package services

// запрос для удаления машины
func QueryCarDelete(id string) string {
	var stroka string = "delete from cars where id = '" + id + "'"
	return stroka
}

// запрос для удаления владельца
func QueryPersonDelete(id string) string {
	var stroka string = "delete from persons where id = '" + id + "'"
	return stroka
}

// запрос для добавления машины
func QueryCarInsert(id string) string {
	var stroka string = ""
	return stroka
}

// запрос для добавления владельца
func QueryPersonInsert(id string) string {
	var stroka string = ""
	return stroka
}
