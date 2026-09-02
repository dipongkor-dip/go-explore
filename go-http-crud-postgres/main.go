package main

import (
	"context"
	"fmt"
	"go-http-crud-postgres/handlers"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	var err = godotenv.Load()
	if err != nil {
		panic("env file not found")
	}

	connectDb()
	defer db.Close(context.Background())

	// Create an instance of PostgresDriver and pass the global db connection
	var driver = &handlers.PostgresDriver{
		DB: db,
	}

	fmt.Println("Server is running at port 5000")

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to go server!")
	})
	http.HandleFunc("POST /users", driver.CreateUserHandler)
	http.HandleFunc("GET /users", driver.GetUsersHandler)
	http.HandleFunc("PUT /users/{id}", driver.UpdateUserHandler)
	http.HandleFunc("DELETE /users/{id}", driver.DeleteUserHandler)
	http.HandleFunc("GET /users/{id}", driver.GetSingleUserHandler)

	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
