package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	o, _ := models.GetOperator(1)

	o.Name = "hello"
	o.Email = "hello@hello.com"
	o.UpdateOperator()
	o, _ = models.GetOperator(1)
	fmt.Println(o)
}
