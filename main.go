package main

import "demo_project/router"

func main() {
	r := router.SetUpRouter()

	r.Run(":8080")
}
