package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	o, _ := models.GetOperator(1)
	fmt.Println(o)
}
