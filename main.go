package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	o := &models.Operator{}
	o.Name = "operator1"
	o.Email = "email@email.com"
	o.PassWord = "password"

	fmt.Println(o)

	o.CreateOperator()

	operator, _ := models.GetOperator(1)
	operator.CreateSchool("first school")

	school, _ := models.GetSchool(1)
	school.CreateClass("first class")
}
