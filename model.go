package main

type Usuario struct {
	Edad   int     `json:"edad"`
	Altura float64 `json:"altura"`
	Peso   float64 `json:"peso"`
}

type Ejercicio struct {
	Nombre       string `json:"nombre"`
	Categoria    string `json:"categoria"`
	Repeticiones int    `json:"repeticiones"`
	Series       int    `json:"series"`
}
