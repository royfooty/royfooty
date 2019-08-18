package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

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

	//Creating new file for writing

	file, err := os.Create("C2ImportFamRelSample1.csv")
	if err != nil {
		fmt.Println("Cannot create file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Picking up data from Database
	type rel struct {
		childid    int
		childname  string
		parentname string
	}

	rows, err := db.Query(`SELECT childid,childname,parentname FROM rel`)
	if err != nil {
		fmt.Println("query problem", err)
	}
	defer rows.Close()

	data := [][]string{} //2D array to store the values(only strings)

	for rows.Next() {
		relt := rel{}
		err = rows.Scan(&relt.childid, &relt.childname, &relt.parentname)
		if err != nil {
			panic(err)
		}
		a := strconv.Itoa(relt.childid)
		data = append(data, []string{a, relt.childname, relt.parentname})
	}

	//Writing data into the csv file
	for _, value := range data {

		err := writer.Write(value)
		if err != nil {
			fmt.Println("Cannot write to file", err)
		}

	}
	fmt.Println("Write Successful")

}
