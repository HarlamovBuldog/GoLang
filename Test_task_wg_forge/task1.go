package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type CatsStat struct {
	color string
	count int64
}

func main() {

	connStr := "host=10.10.0.89 port=5432 user=wg_forge password=42a dbname=wg_forge_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT color, count(*) FROM cats GROUP BY color LIMIT 1000")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		testCatsStat := new(CatsStat)
		if err := rows.Scan(&testCatsStat.color, &testCatsStat.count); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(testCatsStat)

		result, err := db.Exec("insert into cat_colors_info values ($1, $2)",
			testCatsStat.color, testCatsStat.count)

		if err != nil {
			panic(err)
		}

		fmt.Println(result.RowsAffected())
	}
}
