package main

import (
	"demo_project/initializers"
	"demo_project/router"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.SyncDatabase()
}

func main() {

	r := router.SetUpRouter()

	r.Run(":8080")
}
