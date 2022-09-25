package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	school, _ := models.GetSchool(3)
	school.CreateClub("third club", "third content")

	clubs, _ := models.GetClubs()
	for _, v := range clubs {
		fmt.Println(v)
	}
}
