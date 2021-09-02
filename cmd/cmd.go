package cmd

import (
	config "crud-api-homework/configs"
	"crud-api-homework/routes"
	"fmt"
)

func Run() {
	fmt.Println("Masuk cmd")
	config.InitDB()
	routes.Routes()
}
