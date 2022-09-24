package main

import (
	"fmt"
	"log"

	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)
	// controllers.StartmainServer()

	operator, _ := models.GetOperatorByEmail("email@email.com")
	fmt.Println(operator)

	session, err := operator.CreateSession()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(session)
}
