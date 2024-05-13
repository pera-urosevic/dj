package sync

import (
	"os"
	"strings"
)

func getFiles(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	files := []string{}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		files = append(files, entry.Name())
	}
	return files
}
