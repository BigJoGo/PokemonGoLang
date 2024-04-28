package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Импортируем драйвер SQLite3
)

func main() {
	// Открываем соединение с базой данных SQLite3
	db, err := sql.Open("sqlite3", "./pokemon1.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pokemon1 (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT,
		Type1 TEXT,
		Type2 TEXT,
		CP_Per_Level REAL,
		Max_CP INTEGER,
		Max_HP INTEGER,
		Stamina INTEGER,
		Attack INTEGER,
		Defense INTEGER,
		Catch_Percentage REAL,
		Flee_Percentage REAL,
		Current_Form INTEGER,
		Next_Form INTEGER,
		Candy_Required INTEGER
	)`)
	if err != nil {
		fmt.Println("Error create table", err)
		os.Exit(1)
	}
}
