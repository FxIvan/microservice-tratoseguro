package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/api/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		// Lógica específica de tu API Gateway antes de enviar la solicitud al microservicio auth/signup
		fmt.Println("Manejando solicitud en el API Gateway para /auth/signup")

		// Envia la solicitud al microservicio auth/signup en tu Docker
		// Puedes utilizar un cliente HTTP para enviar la solicitud al microservicio
		resp, err := http.Post("http://localhost:4000/auth/signup", "application/json", r.Body)
		if err != nil {
			http.Error(w, "Error al enviar la solicitud al microservicio", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Copia las cabeceras del microservicio a la respuesta del cliente
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		// Copia el código de estado del microservicio a la respuesta del cliente
		w.WriteHeader(resp.StatusCode)

		// Copia el cuerpo de la respuesta del microservicio al cuerpo de la respuesta del cliente
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			fmt.Println("Error al copiar el cuerpo de la respuesta:", err)
		}
	})

	// Inicia el servidor usando http.ListenAndServe
	err := http.ListenAndServe(":8090", router)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
