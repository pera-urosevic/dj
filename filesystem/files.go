package filesystem

import (
	"os"
	"path/filepath"
	"strings"

	"somnusalis.org/dj/database"
)

func GetFiles(path string) []database.RecordSong {
	files := []database.RecordSong{}

	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subdirPath := filepath.Join(path, entry.Name())
			subdirFiles := GetFiles(subdirPath)
			files = append(files, subdirFiles...)
		}

		if !strings.HasSuffix(entry.Name(), ".mp3") {
			continue
		}

		p := filepath.Join(path, entry.Name())
		stats, err := os.Stat(p)
		if err != nil {
			panic(err)
		}

		path := p
		datetime := stats.ModTime()
		file := database.RecordSong{Path: path, Datetime: datetime}
		files = append(files, file)
	}

	return files
}
