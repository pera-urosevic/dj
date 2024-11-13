package playlist

import (
	"os"
	"strings"

	"somnusalis.org/dj/database"
	"somnusalis.org/dj/log"
)

func Playlists() {
	log.Header("PLAYLISTS")

	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `queries`")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var name string
		var query string
		err := rows.Scan(&name, &query)
		if err != nil {
			panic(err)
		}
		log.Action("playlist", name)

		sql := database.GetQuery(db, query)

		rows, err := db.Query(sql)
		if err != nil {
			panic(err)
		}

		lines := []string{}
		lines = append(lines, "#EXTM3U")
		lines = append(lines, "")
		for rows.Next() {
			var path string
			err := rows.Scan(&path)
			if err != nil {
				panic(err)
			}
			log.Action("adding", path)
			lines = append(lines, path)
			lines = append(lines, "")
		}
		m3u := strings.Join(lines, "\n")
		err = os.WriteFile(name+".m3u", []byte(m3u), 0644)
		if err != nil {
			panic(err)
		}

	}

	log.Footer("DONE")
}
