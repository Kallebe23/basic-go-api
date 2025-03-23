package config

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

const db_filepath = "../foo.db"

func BootstrapDB() {
	err := os.WriteFile(db_filepath, []byte{}, 0666)
	if err != nil {
		log.Fatal(err)
	}

	DB, err = sql.Open("sqlite", db_filepath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`create table products (name varchar(255), description text, price float)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`insert into products (name, description, price) values ('camiseta', 'vermelha tamanho G', 22.50)`)
	if err != nil {
		log.Fatal(err)
	}
}