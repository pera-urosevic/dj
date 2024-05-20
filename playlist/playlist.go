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

	where := database.GetQuery(db, query)

	rows, err := db.Query("SELECT songs.path, meta.key, meta.value FROM songs, json_each(songs.meta) as meta WHERE " + where)
	if err != nil {
		panic(err)
	}

	lines := []string{}
	lines = append(lines, "#EXTM3U")
	lines = append(lines, "")
	for rows.Next() {
		result := database.RecordQueryResult{}
		err := rows.Scan(&result.Path, &result.Key, &result.Value)
		if err != nil {
			panic(err)
		}
		log.Action("adding", result.Path)
		lines = append(lines, result.Path)
		lines = append(lines, "")
	}
	m3u := strings.Join(lines, "\n")
	err = os.WriteFile(playlistPath, []byte(m3u), 0644)
	if err != nil {
		panic(err)
	}

	log.Footer("DONE")
}
