package main

import (
	"log"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/handlers"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
	}

	handlers.Manejador()
}
