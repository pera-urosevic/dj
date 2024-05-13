package sync

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"
)

func getMapping(playlist []string) map[string]string {
	mapping := map[string]string{}
	for _, path := range playlist {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		meta, err := tag.ReadFrom(file)
		if err != nil {
			panic(err)
		}
		artist := meta.Artist()
		year := fmt.Sprintf("%d", meta.Year())
		album := meta.Album()
		trackNumber, _ := meta.Track()
		track := fmt.Sprintf("%02d", trackNumber)
		title := meta.Title()
		ext := filepath.Ext(path)
		mapping[path] = sanitize(artist + " (" + year + ") " + album + " - " + track + ". " + title + ext)
	}
	return mapping
}
