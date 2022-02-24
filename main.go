package main

import (
	"github.com/AndreCDiniz/api-go-gin/database"
	"github.com/AndreCDiniz/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
