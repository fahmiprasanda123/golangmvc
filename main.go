package main

import (
	"crud_api_mvc/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
