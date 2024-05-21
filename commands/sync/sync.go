package sync

import (
	"somnusalis.org/dj/database"
	"somnusalis.org/dj/filesystem"
	"somnusalis.org/dj/log"
)

func Sync(musicPath string) {
	log.Header("SYNC")

	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	files := filesystem.GetFiles(musicPath)
	records := getRecords(db)
	records = removeRecords(db, records, files)
	records = addRecords(db, records, files)

	if len(records) == 0 {
		log.Warning("empty database")
	}

	log.Footer("DONE")
}
