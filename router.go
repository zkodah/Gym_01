package main

import "net/http"

func registrarRutas() {
	// Ruta raíz (opcional, solo para probar)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("AppGym corriendo"))
	})

	// Rutas de usuario
	http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			obtenerUsuario(w, r)
		} else if r.Method == http.MethodPost {
			guardarUsuario(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// Ruta de ejercicios
	http.HandleFunc("/ejercicios", listarEjercicios)
}
