package sync

import (
	"database/sql"
	"encoding/json"

	"somnusalis.org/dj/database"
	"somnusalis.org/dj/filesystem"
	"somnusalis.org/dj/log"
)

func getRecords(db *sql.DB) []database.RecordSong {
	rows, err := db.Query("SELECT path, datetime FROM songs")
	if err != nil {
		panic(err)
	}

	records := []database.RecordSong{}
	for rows.Next() {
		record := database.RecordSong{}
		err := rows.Scan(&record.Path, &record.Datetime)
		if err != nil {
			panic(err)
		}
		records = append(records, record)
	}

	return records
}

func removeRecords(db *sql.DB, records []database.RecordSong, files []database.RecordSong) []database.RecordSong {
	foundRecords := []database.RecordSong{}
	for _, record := range records {
		found := false
		for _, file := range files {
			if record.Path == file.Path && record.Datetime == file.Datetime {
				found = true
				break
			}
		}
		if found {
			foundRecords = append(foundRecords, record)
		} else {
			log.Action("removing record", record.Path)
			_, err := db.Exec("DELETE FROM songs WHERE path = ?", record.Path)
			if err != nil {
				panic(err)
			}
		}
	}
	return foundRecords
}

func addRecords(db *sql.DB, records []database.RecordSong, files []database.RecordSong) []database.RecordSong {
	for _, file := range files {
		found := false
		for _, record := range records {
			if record.Path == file.Path {
				found = true
				break
			}
		}
		if !found {
			log.Action("adding record", file.Path)
			file.Meta = filesystem.ReadMeta(file.Path)
			jsonMeta, err := json.Marshal(file.Meta)
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("INSERT INTO songs (path, meta, datetime) VALUES (?, ?, ?)", file.Path, string(jsonMeta), file.Datetime)
			if err != nil {
				panic(err)
			}
			records = append(records, file)
		}
	}
	return records
}
