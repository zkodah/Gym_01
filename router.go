package main

import "net/http"

func habilitarCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}
		next(w, r)
	}
}

func registrarRutas() {
	// Ruta raíz
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(" API de AppGym corriendo"))
	})

	// Usuario
	http.HandleFunc("/usuario", habilitarCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			obtenerUsuario(w, r)
		} else if r.Method == http.MethodPost {
			guardarUsuario(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}))

	// Ejercicios
	http.HandleFunc("/ejercicios", habilitarCORS(listarEjercicios))
}
