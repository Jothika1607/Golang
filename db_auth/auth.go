package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	// "golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// User structure for database interactions
type User struct {
	ID       int
	Username string
	Password string
}

var db *sql.DB

func initDB() {
	config := mysql.Config{
		User:   "jo",
		Passwd: "password",
		Addr:   "localhost:3306",
		Net:    "tcp",
		DBName: "auth_db",
	}

	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")

	http.Handle("/", r)
	fmt.Println("Server listening on port:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	err :=r.ParseForm()

	username := r.FormValue("username")
	password := r.FormValue("password")

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	http.Error(w, "Failed to hash password", http.StatusInternalServerError)
	// 	return
	// }

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	
	fmt.Fprintf(w, "User registered successfully")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Failed to authenticate user", http.StatusInternalServerError)
		return
	}

	// err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if storedPassword != password{

		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	// if err != nil {

	// }

	fmt.Fprintf(w, "Login successful")
}
