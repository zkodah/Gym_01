package main

import (
	"encoding/json"
	"net/http"
	"os"
)

var usuario Usuario

// Guardar usuario
func guardarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "Error en el formato del JSON", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Usuario guardado correctamente"})
}

// Obtener usuario
func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuario)
}

// Listar ejercicios desde JSON
func listarEjercicios(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("data/ejercicios.json")
	if err != nil {
		http.Error(w, "No se pudo abrir ejercicios.json", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var ejercicios []Ejercicio
	if err := json.NewDecoder(file).Decode(&ejercicios); err != nil {
		http.Error(w, "Error al leer ejercicios", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ejercicios)
}
