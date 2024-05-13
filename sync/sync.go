package sync

import (
	"fmt"
)

func Sync(playlistPath string, destinationPath string) {
	fmt.Println()
	fmt.Println("SYNC", playlistPath, destinationPath)
	fmt.Println()

	playlist := getPlaylist(playlistPath)
	mapping := getMapping(playlist)
	files := getFiles(destinationPath)

	// delete files that are not in the playlist
	for _, file := range files {
		found := false
		for _, v := range mapping {
			if v == file {
				found = true
				break
			}
		}
		if !found {
			delete(destinationPath, file)
		}
	}

	// copy files that are not in the destination
	for path, filename := range mapping {
		if exists(destinationPath, filename) {
			continue
		}
		copy(path, destinationPath, filename)
	}

	fmt.Println("DONE")
	fmt.Println()
}
