package services

import (
	"module/internal/models"
	"strconv"
)

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

func ConstructWhere(current models.Car) string {

	var result string = "where "
	if current.Id != "" {
		result += "car.id = '" + current.Id + "' and "
	}
	if current.Mark != "" {
		result += "mark = '" + current.Mark + "' and "
	}
	if current.Model != "" {
		result += "model = '" + current.Model + "' and "
	}
	if current.RegNum != "" {
		result += "regNum = '" + current.RegNum + "' and "
	}
	if current.Year != "" {
		result += "year = '" + current.Year + "' and "
	}
	if current.Owner.Name != "" {
		result += "name = '" + current.Owner.Name + "' and "
	}
	if current.Owner.Surname != "" {
		result += "surname = '" + current.Owner.Surname + "' and "
	}
	if current.Owner.Patronymic != "" {
		result += "patronymic = '" + current.Owner.Patronymic + "' and "
	}

	newRes := result[:len(result)-4]
	return newRes
}

func ConStringShowSpec(car *models.Car, limit string, offset string) string {

	var stroka string = "SELECT car.id,regnum,mark,model,year,people.id,name,surname,patronymic FROM car left join people on car.owner = people.id "

	newRes := stroka + ConstructWhere(*car)

	test, _ := strconv.Atoi(limit)
	if test > 0 {
		newRes += "limit " + limit + " "
	}
	test, _ = strconv.Atoi(offset)
	if test > 0 {
		newRes += "offset " + offset + " "
	}

	return newRes
}
