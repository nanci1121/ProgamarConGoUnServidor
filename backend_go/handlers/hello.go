package handlers

import (
    "fmt"
    "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola, mundo desde Go soy Venancio!")
}