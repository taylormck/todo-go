package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
		create table if not exists todos (
			id integer primary key autoincrement,
			title text,
			status text
		);
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateTodo(title string, status string) (int64, error) {
	result, err := DB.Exec("insert into todos (title, status) values (?, ?)", title, status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteTodo(id int64) error {
	_, err := DB.Exec("delete from todos where id = ?", id)
	return err
}

func ReadTodoList() []Todo {
	rows, _ := DB.Query("select id, title, status from todos")
	defer rows.Close()

	todos := make([]Todo, 0)
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Status)
		todos = append(todos, todo)
	}
	return todos
}
