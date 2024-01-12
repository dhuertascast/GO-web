package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Manejador para el endpoint /ping
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// Verificar que el método de solicitud sea GET
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Responder con el texto "pong"
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "pong")
	})

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
