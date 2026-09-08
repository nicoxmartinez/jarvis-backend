package main

import (
	"fmt"
	"jarvis-backend/routers"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	routers.SetServer()
	fmt.Printf("Servidor J.A.R.V.I.S. escuchando en http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}
