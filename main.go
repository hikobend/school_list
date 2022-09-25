package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	school2, _ := models.GetSchool(2)
	clubs, _ := school2.GetClubBySchool()
	for _, v := range clubs {
		fmt.Println(v)
	}
}
