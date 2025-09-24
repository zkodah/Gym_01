package main

type Usuario struct {
	Edad   int     `json:"edad"`
	Altura float64 `json:"altura"`
	Peso   float64 `json:"peso"`
}

// Estructura correcta para tu JSON de ejercicios
type Ejercicio struct {
	Dia    string   `json:"dia"`
	Titulo string   `json:"titulo"`
	Rutina []string `json:"rutina"`
}
