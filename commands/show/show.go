package show

import (
	"facette.io/natsort"
	"golang.org/x/exp/maps"
	"somnusalis.org/dj/database"
	"somnusalis.org/dj/log"
)

func Show(path string) {
	log.Header("SHOW", path)

	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT key, value FROM songs, json_each(songs.meta) WHERE songs.path LIKE ?", "%"+path+"%")
	if err != nil {
		panic(err)
	}

	meta := map[string]interface{}{}
	for rows.Next() {
		result := database.RecordQueryResult{}
		err := rows.Scan(&result.Key, &result.Value)
		if err != nil {
			panic(err)
		}
		meta[result.Key] = result.Value
	}
	keys := maps.Keys(meta)
	natsort.Sort(keys)
	for _, k := range keys {
		v := meta[k]
		log.Action(k, v)
	}

	log.Footer("DONE")
}
