package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)

	c, _ := models.GetClub(1)
	c.DeleteClub()
}
