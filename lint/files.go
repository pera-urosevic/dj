package lint

import (
	"os"
	"path/filepath"
	"strings"
)

func listFiles(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	files := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			subPath := filepath.Join(path, entry.Name())
			subFiles := listFiles(subPath)
			files = append(files, subFiles...)
			continue
		}
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		filePath := filepath.Join(path, entry.Name())
		files = append(files, filePath)
	}
	return files
}
