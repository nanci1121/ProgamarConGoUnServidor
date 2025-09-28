package handlers

import (
    "database/sql"
    "fmt"
    "net/http"
    "os"

    "golang.org/x/crypto/bcrypt"
    _ "github.com/lib/pq"
)



// Handler para login de usuario
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Cargar variables de entorno
    loadEnv()

    // Obtener datos del formulario
    email := r.FormValue("email")
    password := r.FormValue("password")

    // Construir la cadena de conexión usando las variables de entorno
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    // Conectar a la base de datos
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        http.Error(w, "Error de conexión a la base de datos", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Buscar el hash de la contraseña del usuario por email
    var hashedPassword string
    err = db.QueryRow("SELECT password FROM usuarios WHERE email=$1", email).Scan(&hashedPassword)
    if err != nil {
        http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
        return
    }

    // Comparar la contraseña ingresada con el hash almacenado
    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
        return
    }

    fmt.Fprintf(w, "Login exitoso")
}