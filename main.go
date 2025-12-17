package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database connection
var db *sql.DB

// User model
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Response wrapper
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database connection
	initDB()
	defer db.Close()

	// Initialize router
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/health", healthCheckHandler).Methods("GET")
	api.HandleFunc("/users", getUsersHandler).Methods("GET")
	api.HandleFunc("/users", createUserHandler).Methods("POST")
	api.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	api.HandleFunc("/users/{id}", updateUserHandler).Methods("PUT")
	api.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")

	// Serve static files for frontend
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend")))

	// CORS middleware
	router.Use(corsMiddleware)

	// Start server
	serverPort := getEnv("SERVER_PORT", "8080")
	serverHost := getEnv("SERVER_HOST", "0.0.0.0")
	addr := fmt.Sprintf("%s:%s", serverHost, serverPort)

	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

func initDB() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "gogo_db")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test connection
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connection established")

	// Run migrations
	runMigrations()
}

func runMigrations() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	log.Println("Migrations completed successfully")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// In production, this should be restricted to specific origins
		allowedOrigin := getEnv("CORS_ALLOWED_ORIGIN", "*")
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Message: "Server is running",
		Data: map[string]string{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
		},
	}
	respondJSON(w, http.StatusOK, response)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email, created_at FROM users ORDER BY created_at DESC")
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			respondError(w, http.StatusInternalServerError, "Failed to scan user")
			return
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		respondError(w, http.StatusInternalServerError, "Error iterating over users")
		return
	}

	response := Response{
		Success: true,
		Data:    users,
	}
	respondJSON(w, http.StatusOK, response)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	err := db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err == sql.ErrNoRows {
		respondError(w, http.StatusNotFound, "User not found")
		return
	} else if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	response := Response{
		Success: true,
		Data:    user,
	}
	respondJSON(w, http.StatusOK, response)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if user.Name == "" || user.Email == "" {
		respondError(w, http.StatusBadRequest, "Name and email are required")
		return
	}

	err := db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at",
		user.Name, user.Email,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	response := Response{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	}
	respondJSON(w, http.StatusCreated, response)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if user.Name == "" || user.Email == "" {
		respondError(w, http.StatusBadRequest, "Name and email are required")
		return
	}

	result, err := db.Exec(
		"UPDATE users SET name = $1, email = $2 WHERE id = $3",
		user.Name, user.Email, id,
	)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to get rows affected")
		return
	}
	if rowsAffected == 0 {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	response := Response{
		Success: true,
		Message: "User updated successfully",
	}
	respondJSON(w, http.StatusOK, response)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	result, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to get rows affected")
		return
	}
	if rowsAffected == 0 {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	response := Response{
		Success: true,
		Message: "User deleted successfully",
	}
	respondJSON(w, http.StatusOK, response)
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	response := Response{
		Success: false,
		Message: message,
	}
	respondJSON(w, status, response)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
