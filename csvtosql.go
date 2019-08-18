package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Cronaldo@7"
	dbname   = "test2"
)

func main() {

	//connecting to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err1 := sql.Open("postgres", psqlInfo)

	if err1 != nil {
		panic(err1)
	}

	defer db.Close()

	fmt.Println("Successfully connected to database")

	//opening a csv file to read its contents
	file, err := os.Open("C2ImportFamRelSample.csv")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(record)

	db.Query(`TRUNCATE REL`) //EVERY TIME A NEW RECORD IS PUT,THE TABLE HAS TO BE UPDATED

	stmt := `INSERT INTO REL (childid,childname,parentname) VALUES($1,$2,$3);`
	for _, row := range record[0:] {
		_, err := db.Query(stmt, row[0], row[1], row[2])
		if err != nil {
			log.Fatal(err)
		}
	}

}
