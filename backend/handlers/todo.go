// CRUDのハンドラ作成
package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	models "github.com/koheiterajima-bs/restapi-mysql-golang-docker-hands-on01/backend/models"
)

var db *sql.DB

func init() {
	var err error
	dsn := "todo_user:todo_password@tcp(db:3306)/todo_app"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func HandleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTodos(w, r)
	case http.MethodPost:
		CreateTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	if id == "" {
		http.Error(w, "Todo ID is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		UpdateTodo(w, r, id)
	case http.MethodDelete:
		DeleteTodo(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetTodos(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.todo
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &todo)

	if err := models.CreateTodo(db, &todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request, id string) {
	var todo models.todo
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &todo)

	if err := models.UpdateTodo(db, id, &todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request, id string) {
	if err := models.DeleteTodo(db, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
