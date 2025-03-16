package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"forum/Handlers"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	// Initialisation de la base de données
	InitDB()

	// Définition des routes

	// Login
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.LoginFromHandler(w, r)
		} else if r.Method == http.MethodPost {
			handlers.LoginHandler(w, r, db)
		} else {
			http.Error(w, "main.go - /login", http.StatusMethodNotAllowed)
		}
	})

	// Create Posts 
	http.HandleFunc("/create-post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreatePostHandler(w, r, db)
		} else {
			http.Error(w, "main.go - /create-post", http.StatusMethodNotAllowed)
		}
	})

	// Route pour afficher le formulaire de création de post (nouveau handler)
	http.HandleFunc("/new-post", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/new_post.html")
	})

	// Static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Lancement du serveur
	port := ":8081"
	fmt.Println("Serveur démarré sur http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// InitDB initialise la connexion à la base de données
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`

	postsTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Erreur lors de la création des tables: ", err)
	}

	_, err = db.Exec(postsTable)
	if err != nil {
		log.Fatal("Erreur de la création de la table posts", err)
	}
}