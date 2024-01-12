package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura para el cuerpo de la solicitud JSON
type GreetingRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	// Manejador para el endpoint /greetings
	http.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
		// Verificar que el método de solicitud sea POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Decodificar el cuerpo de la solicitud JSON
		var request GreetingRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Construir la respuesta
		response := fmt.Sprintf("Hello %s %s", request.FirstName, request.LastName)

		// Responder con el saludo
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, response)
	})

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
