package main

import (
	"net/http"
)

func registrarRutas() {
	http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			guardarUsuario(w, r)
		} else if r.Method == http.MethodGet {
			obtenerUsuario(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/ejercicios", listarEjercicios)
}
