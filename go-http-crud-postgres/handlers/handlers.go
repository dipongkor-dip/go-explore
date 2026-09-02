package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-http-crud-postgres/models"

	"github.com/jackc/pgx/v5"
)

type PostgresDriver struct {
	DB *pgx.Conn
}

func (d *PostgresDriver) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	var err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	var query = `
		INSERT INTO users (username, age, email) VALUES($1, $2, $3)
		returning id
	`
	err = d.DB.QueryRow(context.Background(), query, newUser.Name, newUser.Age, newUser.Email).Scan(&newUser.Id)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not create user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (d *PostgresDriver) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	var query = `
		SELECT id, username, age, email FROM users
	`
	var rows, err = d.DB.Query(context.Background(), query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not get users")
		return
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		var err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Could not scan user")
			return
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not read user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(users)
}

func (d *PostgresDriver) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var idParam = r.PathValue("id")

	var id, err = strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	var updateUser models.User

	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	var query = `
	UPDATE users SET username = $1, age = $2, email = $3 WHERE id = $4
	returning id, username, age, email
	`
	err = d.DB.QueryRow(context.Background(), query, updateUser.Name, updateUser.Age, updateUser.Email, id).Scan(&updateUser.Id, &updateUser.Name, &updateUser.Age, &updateUser.Email)

	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could no update user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateUser)
}

func (d *PostgresDriver) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var idParam = r.PathValue("id")

	var id, err = strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	var query = `DELETE FROM users WHERE id = $1`

	var cmdTag, er = d.DB.Exec(context.Background(), query, id)

	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not delete user")
		return
	}

	if cmdTag.RowsAffected() != 1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "User deleted successfully")
}

func (d *PostgresDriver) GetSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	var idParam = r.PathValue("id")

	var id, err = strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	var user models.User

	var query = `SELECT id, username, age, email FROM users WHERE id = $1`

	err = d.DB.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Name, &user.Age, &user.Email)

	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not get user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(user)
}
