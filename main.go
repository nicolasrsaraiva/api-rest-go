package main

import (
	"fmt"

	"github.com/nicolasrsaraiva/api-rest-go/database"
	"github.com/nicolasrsaraiva/api-rest-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor Rest com go")
	routes.HandleRequest()
}
