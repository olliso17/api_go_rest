package main

import (
	"api_go_rest/database"
	"api_go_rest/models"
	"api_go_rest/routes"
	"fmt"
)

// no arquivo main vc executa as outras funções
func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDadoss()
	fmt.Println("Iniciando api rest com go")
	routes.HandleRequest()
}
