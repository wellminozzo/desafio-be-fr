package main

import (
	"github.com/labstack/echo"
	"github.com/wellminozzo/desafio-be-fr/cmd"
	"github.com/wellminozzo/desafio-be-fr/freterapido"
	"github.com/wellminozzo/desafio-be-fr/models"
)

func main() {
	models.InitDB() // Função para inicializar o banco de dados
	e := echo.New()

	freterapido.Routes(e)

	cmd.Execute()

}
