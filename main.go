package main

import (
	"fmt"
	"jarvis-backend/routers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando variables de entorno")
	}
	port := os.Getenv("PORT")

	routers.SetServer()
	fmt.Printf("Servidor J.A.R.V.I.S. escuchando en http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}
