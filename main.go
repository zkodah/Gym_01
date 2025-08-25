package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")

	registrarRutas()

	http.ListenAndServe(":8080", nil)
}
