package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

// Handler pour la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Accès à la page d'accueil..")
	tmpl, err := template.ParseFiles("templates/accueil.html")
	if err != nil {
		http.Error(w, "Erreur interne :", http.StatusInternalServerError)
		fmt.Println("Erreur lors du chargement de la template :", err)
		return
	}
	tmpl.Execute(w, nil)
}

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

func LoginFromHandler(w http.ResponseWriter, r * http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Erreur Interne :", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}