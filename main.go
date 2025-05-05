package main

import (
	"fmt"
	"net/http"
)

type UVData struct {
	NivelUV       float64 `json:"nivel_uv"`
	Classificacao string  `json:"classificacao"`
	Timestamp     string  `json:"timestamp"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sensor UV ativo na porta 8080!")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status Ok!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Servidor rodando na porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
