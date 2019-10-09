package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "docker-user"
	dbPassword = "docker-password"
	dbName     = "docker-dbname"
	dbPort     = "5432"
)

type person struct {
	ID      int
	name    string
	surname string
	age     int
}

func main() {
	// dbInfo := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", dbUser, dbPassword, dbPort, dbName)
	dbInfo := fmt.Sprintf("dbname=%s user=%s password=%s host=0.0.0.0 port=%s sslmode=disable", dbName, dbUser, dbPassword, dbPort)
	// fmt.Println(dbInfo)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("# Inserting values")
	db.Exec("CREATE TABLE person (id int PRIMARY KEY NOT NULL, name varchar(80), surname varchar(80), age int);")

	i := 0

	for {
		var p person
		p.ID = i

		print("Input name: ")
		_, err = fmt.Scanln(&p.name)
		if err != nil {
			log.Fatal(err)
		}
		p.name = strings.TrimSpace(p.name)

		print("Input surname: ")
		_, err = fmt.Scanln(&p.surname)
		if err != nil {
			log.Fatal(err)
		}
		p.surname = strings.TrimSpace(p.surname)

		print("Input age: ")
		_, err := fmt.Scan(&p.age)
		if err != nil {
			log.Fatal(err)
		}

		query := fmt.Sprintf("INSERT INTO person(id, name, surname, age) VALUES(%d, '%s', '%s', %d)", p.ID, p.name, p.surname, p.age)

		// fmt.Println(p)
		// fmt.Println(query)

		db.Exec(query)

	}
}
