package query

import (
	"somnusalis.org/dj/database"
	"somnusalis.org/dj/log"
)

func Queries() {
	log.Header("QUERIES")

	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, query FROM queries")
	if err != nil {
		panic(err)
	}

	i := 0
	for rows.Next() {
		result := database.RecordQuery{}
		err := rows.Scan(&result.Name, &result.Query)
		if err != nil {
			panic(err)
		}
		i++
		log.KeyValue(result.Name, result.Query)
	}

	log.Footer("DONE")
}
