package sync

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copy(source string, destinationPath string, destinationFilename string) {
	destination := filepath.Join(destinationPath, destinationFilename)
	fmt.Println("COPY")
	fmt.Println(source)
	fmt.Println(destination)
	fmt.Println()
	s, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	d, err := os.Create(destination)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		panic(err)
	}
}
