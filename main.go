package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	school, _ := models.GetSchool(2)
	school.CreateClub("first Club", "first content")
}
