package main

import (
	"eshop-backend-go/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
