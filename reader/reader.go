package main

import (
	"database/sql"
	"fmt"
	"log"

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
	dbInfo := fmt.Sprintf("dbname=%s user=%s password=%s host=0.0.0.0 port=%s sslmode=disable", dbName, dbUser, dbPassword, dbPort)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, _ := db.Query("SELECT * FROM person")
	for rows.Next() {
		p := person{}
		err := rows.Scan(&p.ID, &p.name, &p.surname, &p.age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(p)
	}
}
