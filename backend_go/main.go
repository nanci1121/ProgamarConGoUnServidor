package main

import (
    "fmt"
    "net/http"
    "ProgramarGo/handlers"
)

func main() {
    http.HandleFunc("/", handlers.HelloHandler)
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/register", handlers.RegisterHandler)
    
    fmt.Println("Servidor seguro escuchando en https://localhost:8443")
    err := http.ListenAndServeTLS(":8443", "certs/server.crt", "certs/server.key", nil)
    if err != nil {
        fmt.Println("Error iniciando el servidor HTTPS:", err)
    }
}