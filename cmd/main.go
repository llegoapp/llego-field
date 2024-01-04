package main

import (
	"fields/pkg/database"
	"fmt"
)

func main() {

	fmt.Println("Hello World!")

	err := database.InitPool("app/migrations")
	if err != nil {
		fmt.Printf("error initializing database: %v\n", err)
		return
	}
	println("database initialized successfully")

}
