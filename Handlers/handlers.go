package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"
)


// Handler pour l'inscription
func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, password)
	if err != nil {
		http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Inscription réussie !")
}

// Handler pour la connexion
func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	var dbPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&dbPassword)
	if err != nil {
		http.Error(w, "Identifiants incorrects", http.StatusUnauthorized)
		return
	}

	if password != dbPassword {
		http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Connexion réussie !")
}

func LoginFromHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/log&Singup.html")
	if err != nil {
		http.Error(w, "handlers.go 60 :", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Handler des posts
func CreatePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "CreatePost - handlers - 81", http.StatusMethodNotAllowed)
		return
	}

	userID := 1
	title := r.FormValue("title")
	content := r.FormValue("content")

	if title == " " || content == " " {
		http.Error(w, "CreatePost - Handlers - 90", http.StatusMethodNotAllowed)
		return
	}

	_, err := db.Exec("INSERT INTO posts (user_id, title, content, created_at) VALUES (?, ?, ?, ?)", userID, title, content, time.Now())
	if err != nil {
		http.Error(w, "CreatePost - Handlers - 96", http.StatusInternalServerError)
		fmt.Println("Erreur SQL:", err)
		return
	}
}
