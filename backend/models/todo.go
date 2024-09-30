// データベースとやり取りするためのモデル作成
package models

import "database/sql"

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodos(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func CreateTodo(db *sql.DB, todo *Todo) error {
	result, err := db.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	todo.ID = int(id)
	return nil
}

func UpdateTodo(db *sql.DB, id string, todo *Todo) error {
	_, err := db.Exec("UPDATE todos SET title = ?, completed = ?, WHERE id = ?", todo.Title, todo.Completed, id)
	return err
}

func DeleteTodo(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
