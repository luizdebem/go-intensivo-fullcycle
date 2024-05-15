// Aula 3 - Desenvolvimento de APIs com Golang

// MUX (multiplexer) condensador de acessos.

package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	// rever pprof e testar com a endpoint do fibonacci depois, no windows não deu pra fazer requisições
	_ "net/http/pprof"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", listUsersHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("/cpu", CPUIntensiveEndpoint)
	go http.ListenAndServe(":3000", mux)
	http.ListenAndServe(":6060", nil) // duas threads com dois servers!
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// defer: quando tudo for executado, fecha a conexão.
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// defer: quando tudo for executado, fecha a conexão.
	defer db.Close()

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := db.Exec("INSERT INTO users (id, name, email) VALUES (?, ?, ?)", u.ID, u.Name, u.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func CPUIntensiveEndpoint(w http.ResponseWriter, r *http.Request) {
	result := fib(60)
	w.Write([]byte(strconv.Itoa(result)))
}

// func GenerateLargeString(n int) string {
// 	var buffer bytes.Buffer
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < 100; j++ {
// 			buffer.WriteString(strconv.Itoa(i + j*j))
// 		}
// 	}
// 	return buffer.String()
// }

func GenerateLargeString(n int) string {
	var buffer bytes.Buffer
	buffer.Grow(n * 100)
	for i := 0; i < n; i++ {
		for j := 0; j < 100; j++ {
			buffer.WriteString(strconv.Itoa(i + j*j))
		}
	}
	return buffer.String()
}
