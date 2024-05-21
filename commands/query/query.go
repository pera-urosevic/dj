package query

import (
	"somnusalis.org/dj/database"
	"somnusalis.org/dj/log"
)

func Query(query string) {
	log.Header("QUERY", query)

	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sql := database.GetQuery(db, query)
	log.Action("querying database", sql)

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	i := 0
	for rows.Next() {
		result := database.RecordSong{}
		err := rows.Scan(&result.Path)
		if err != nil {
			panic(err)
		}
		i++
		log.Result(i, result)
	}

	log.Footer("DONE")
}
