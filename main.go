package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/p2gc1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	uh := UserHandler{NewUserRepository(db)}
	// uh := UserHandler{userRepository}
	router.HandleFunc("/users", uh.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", uh.GetUserByIDHandler).Methods("GET")
	router.HandleFunc("/users/{id}", uh.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", uh.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users", uh.GetAllUsersHandler).Methods("GET")

	fmt.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func NewUserRepository(db *sql.DB) {
	panic("unimplemented")
}
