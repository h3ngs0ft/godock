package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cnn, err := sql.Open("mysql", "docker:docker@tcp(192.168.99.100:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	id := 1
	var name string

	if err := cnn.QueryRow("SELECT name FROM test_tb WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name)
}
