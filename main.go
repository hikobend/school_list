package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	c, _ := models.GetClass(1)
	c.ClassNumber = "Update class_number"
	c.UpdateClass()
}
