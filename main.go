package main

import (
	"os"

	"somnusalis.org/dj/inspect"
	"somnusalis.org/dj/log"
	"somnusalis.org/dj/playlist"
	"somnusalis.org/dj/query"
	"somnusalis.org/dj/sync"
)

func usage() {
	log.Print("")
	log.Print("DJ usage:")
	log.Print("dj sync <folder> - sync folder to database")
	log.Print("dj query <query> - query database")
	log.Print("dj inspect <file> - inspect file")
	log.Print("dj playlist <query> <playlist> - create playlist with query")
	log.Print("")
	os.Exit(0)
}

func main() {
	args := os.Args

	if len(args) < 3 {
		usage()
	}

	// sync
	if (args[1] == "sync") && (args[2] != "") {
		sync.Sync(args[2])
		return
	}

	// query
	if (args[1] == "query") && (args[2] != "") {
		query.Query(args[2])
		return
	}

	// inspect
	if (args[1] == "inspect") && (args[2] != "") {
		inspect.Inspect(args[2])
		return
	}

	if len(args) < 4 {
		usage()
	}

	// playlist
	if (args[1] == "playlist") && (args[2] != "") && (args[3] != "") {
		playlist.Playlist(args[2], args[3])
		return
	}

	usage()
}
