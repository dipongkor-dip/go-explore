package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

func main() {
	var mux = http.NewServeMux()

	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("GET /users", getUserHandler)
	mux.HandleFunc("GET /users/{id}", getSingleUserHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	fmt.Println("Server is running at port 5000")

	var err = http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is up and healthy")
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var users = []User{
	{
		Id:    1,
		Name:  "Sunny",
		Age:   25,
		Email: "sunny@gmail.com",
	},
	{
		Id:    2,
		Name:  "Alo",
		Age:   26,
		Email: "alo@gmail.com",
	},
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	fmt.Fprintln(w, "Method not allowed")
	// 	return
	// }

	var newUser User

	var err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	newUser.Id = len(users) + 1
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// var urs, _ = json.Marshal(users)
	// w.Write(urs)
	// alternative
	json.NewEncoder(w).Encode(users) // memory efficient
}

func getSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	var idParam = r.PathValue("id")

	fmt.Printf("the value or id is %v and the type of the id is %T", idParam, idParam)

	var id, err = strconv.Atoi(idParam) // string to integer convert

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	var isFound bool

	for _, user := range users {
		if user.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			isFound = true
			break
		}
	}

	if !isFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found")
	}
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var idParam = r.PathValue("id")

	var id, err = strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	var newUser User

	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	var isFound bool

	for idx, user := range users {
		if user.Id == id {
			newUser.Id = id
			users[idx] = newUser

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			isFound = true
			break
		}
	}

	if !isFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found")
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var idParam = r.PathValue("id")

	var id, err = strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	var isFound bool

	for idx, user := range users {

		if user.Id == id {
			// users = append(users[:idx], users[idx+1:]...) // 0 to index and index + 1 to end values resign
			// alternative
			users = slices.Delete(users, idx, idx+1)

			w.WriteHeader(http.StatusNoContent)
			isFound = true
			return
		}
	}

	if !isFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found")
	}
}
