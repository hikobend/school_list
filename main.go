package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	o := models.Operator{}
	o.Name = "test"
	o.Email = "test@test.com"
	o.PasaWord = "password"

	fmt.Println(o)

	o.CreateOperator()
}
