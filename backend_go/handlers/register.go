package handlers

import (
    "database/sql"
    "fmt"
    "net/http"
    "os"

    "golang.org/x/crypto/bcrypt"
    _ "github.com/lib/pq"
)



// Handler para registrar un nuevo usuario
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Cargar variables de entorno
    loadEnv()

    // Obtener datos del formulario
    nombre := r.FormValue("nombre")
    direccion := r.FormValue("direccion")
    email := r.FormValue("email")
    telefono := r.FormValue("telefono")
    password := r.FormValue("password")

    // Hashear la contraseña usando bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Error al cifrar la contraseña:", err)
        http.Error(w, "Error al cifrar la contraseña", http.StatusInternalServerError)
        return
    }

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
        fmt.Println("Error de conexión a la base de datos:", err)
        http.Error(w, "Error de conexión a la base de datos", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Insertar el usuario en la base de datos
    _, err = db.Exec(
        "INSERT INTO usuarios (nombre, direccion, email, telefono, password) VALUES ($1, $2, $3, $4, $5)",
        nombre, direccion, email, telefono, string(hashedPassword),
    )
    if err != nil {
        fmt.Println("Error al registrar el usuario:", err)
        http.Error(w, "Error al registrar el usuario", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Usuario registrado correctamente")
}