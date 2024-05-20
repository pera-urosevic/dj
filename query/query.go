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

	where := database.GetQuery(db, query)
	log.Action("querying database where", where)

	rows, err := db.Query("SELECT songs.path, meta.key, meta.value FROM songs, json_each(songs.meta) as meta WHERE " + where)
	if err != nil {
		panic(err)
	}

	i := 0
	for rows.Next() {
		result := database.RecordQueryResult{}
		err := rows.Scan(&result.Path, &result.Key, &result.Value)
		if err != nil {
			panic(err)
		}
		i++
		log.Result(i, result)
	}

	log.Footer("DONE")
}
