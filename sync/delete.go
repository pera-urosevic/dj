package sync

import (
	"fmt"
	"os"
	"path/filepath"
)

func delete(destinationPath string, destinationFilename string) {
	destination := filepath.Join(destinationPath, destinationFilename)
	fmt.Println("DELETE")
	fmt.Println(destination)
	fmt.Println()
	err := os.Remove(destination)
	if err != nil {
		panic(err)
	}
}
