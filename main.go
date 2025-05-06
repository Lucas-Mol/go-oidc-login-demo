// main.go
package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(".env file not found")
	}

	initOIDC()

	http.HandleFunc("/auth/login", handleLogin)
	http.HandleFunc("/auth/callback", handleCallback)
	http.HandleFunc("/me", handleMe)

	log.Println("Server running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
