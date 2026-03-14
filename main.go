package main

import (
	"fmt"
	"log"
	"net/http"
)

// A função que responde à requisição web
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! I'm GO app-a!!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor rodando na porta 8080...")
	
	// Inicia o servidor na porta 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}