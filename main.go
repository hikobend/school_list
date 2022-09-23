package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	o := &models.Operator{}
	o.Name = "test3"
	o.Email = "test3@test3.com"
	o.Name = "test3test"
	fmt.Println(o)

	o.CreateOperator()

	operator, _ := models.GetOperator(3)
	operator.CreateSchool("third school")

	operator2, _ := models.GetOperator(2)
	schools, _ := operator2.GetSchoolByOperator()
	for _, v := range schools {
		fmt.Println(v)
	}
}
