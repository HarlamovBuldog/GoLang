package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/montanaflynn/stats"
)

type CatsStat struct {
	mean   float64
	median float64
	mode   []float64
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

	rows, err := db.Query("SELECT tail_length, whiskers_length FROM cats")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tailLengthArray []float64
	var whiskersLengthArray []float64

	for rows.Next() {
		var iterationTailLength float64
		var iterationWhiskersLength float64

		if err := rows.Scan(&iterationTailLength, &iterationWhiskersLength); err != nil {
			fmt.Println(err.Error())
			return
		}

		tailLengthArray = append(tailLengthArray, iterationTailLength)
		whiskersLengthArray = append(whiskersLengthArray, iterationWhiskersLength)
	}

	fmt.Println(tailLengthArray)
	fmt.Println(whiskersLengthArray)

	tailLengthMean, _ := stats.Mean(tailLengthArray)
	tailLengthMedian, _ := stats.Median(tailLengthArray)
	tailLengthMode, _ := stats.Mode(tailLengthArray)

	var tailLengthModeInt []int
	for _, elem := range tailLengthMode {
		tailLengthModeInt = append(tailLengthModeInt, int(elem))
	}

	whiskersLengthMean, _ := stats.Mean(whiskersLengthArray)
	whiskersLengthMedian, _ := stats.Median(whiskersLengthArray)
	whiskersLengthMode, _ := stats.Mode(whiskersLengthArray)

	var whiskersLengthModeInt []int
	for _, elem := range whiskersLengthMode {
		whiskersLengthModeInt = append(whiskersLengthModeInt, int(elem))
	}

	fmt.Println(tailLengthMean, tailLengthMedian, tailLengthModeInt,
		whiskersLengthMean, whiskersLengthMedian, whiskersLengthModeInt)

	result, err := db.Exec("insert into cat_colors_info values ($1, $2, $3, $4, $5, $6)",
		tailLengthMean, tailLengthMedian, tailLengthModeInt,
		whiskersLengthMean, whiskersLengthMedian, whiskersLengthModeInt)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
	fmt.Println("Success!")

}
