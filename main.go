package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)
	// school, _ := models.GetSchool(3)
	// school.CreateClass("third class")

	school2, _ := models.GetSchool(1)
	classes, _ := school2.GetClassBySchool()
	for _, v := range classes {
		fmt.Println(v)
	}
}
