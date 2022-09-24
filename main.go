package main

import (
	"fmt"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)
	// controllers.StartmainServer()

	operator, _ := models.GetOperatorByEmail("email@email.com")
	fmt.Println(operator)
}
