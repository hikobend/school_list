package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	school, _ := models.GetSchool(2)
	school.CreateClass("second class")

	classes, _ := models.GetClasses()
	for _, v := range classes {
		fmt.Println(v)
	}
}
