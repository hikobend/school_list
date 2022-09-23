package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	s, _ := models.GetSchool(2)
	s.Name = "second school"
	s.UpdateSchool()
}
