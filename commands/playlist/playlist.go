package playlist

import (
	"os"
	"strings"

	"somnusalis.org/dj/database"
	"somnusalis.org/dj/log"
)

func Playlist(query string, playlistPath string) {
	log.Header("PLAYLIST")

	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
	err = os.WriteFile(playlistPath, []byte(m3u), 0644)
	if err != nil {
		panic(err)
	}

	log.Footer("DONE")
}
