package sync

import (
	"os"
	"path/filepath"
)

func exists(path string, filename string) bool {
	_, err := os.Stat(filepath.Join(path, filename))
	return err == nil
}
