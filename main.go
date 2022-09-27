package main

import (
	"fmt"

	"example.com/school/app/controllers"
	"example.com/school/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartmainServer()
}
