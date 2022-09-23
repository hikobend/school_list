package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)
	operator, _ := models.GetOperator(2)
	operator.CreateSchool("testtest")

	schools, _ := models.GetSchools()
	for _, v := range schools {
		fmt.Println(v)
	}
}
