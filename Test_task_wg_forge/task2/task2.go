package main

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/montanaflynn/stats"
)

func main() {
	//do not forget to change host and pass
	connStr := "host=localhost port=5432 user=wg_forge password=a42 dbname=wg_forge_db sslmode=disable"
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

	tailLengthMean = float64(int64(tailLengthMean*10+0.5)) / 10
	tailLengthMedian = float64(int64(tailLengthMedian*10+0.5)) / 10

	whiskersLengthMean, _ := stats.Mean(whiskersLengthArray)
	whiskersLengthMedian, _ := stats.Median(whiskersLengthArray)
	whiskersLengthMode, _ := stats.Mode(whiskersLengthArray)

	whiskersLengthMean = float64(int64(whiskersLengthMean*10+0.5)) / 10
	whiskersLengthMedian = float64(int64(whiskersLengthMedian*10+0.5)) / 10

	fmt.Println(tailLengthMean, tailLengthMedian, tailLengthMode,
		whiskersLengthMean, whiskersLengthMedian, whiskersLengthMode)

	result, err := db.Exec("insert into cats_stat values ($1, $2, $3, $4, $5, $6)",
		tailLengthMean, tailLengthMedian, pq.Array(tailLengthMode),
		whiskersLengthMean, whiskersLengthMedian, pq.Array(whiskersLengthMode))

	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
	fmt.Println("Success!")

}
