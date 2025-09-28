package handlers

import (
    "fmt"
    "github.com/joho/godotenv"
)

// Carga las variables de entorno desde el archivo .env
func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error cargando .env:", err)
    }
}
