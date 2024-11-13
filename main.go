package main

import (
	"os"

	"somnusalis.org/dj/commands/inspect"
	"somnusalis.org/dj/commands/playlist"
	"somnusalis.org/dj/commands/query"
	"somnusalis.org/dj/commands/show"
	"somnusalis.org/dj/commands/sync"
	"somnusalis.org/dj/log"
)

func usage() {
	log.Print("")
	log.Print("DJ usage:")
	log.Print("dj inspect <file> - inspect file")
	log.Print("dj sync <folder> - sync folder to database")
	log.Print("dj show <path> - show meta for path")
	log.Print("dj queries - list queries in database")
	log.Print("dj query <query> - query database")
	log.Print("dj playlists - create playlists for each query")
	log.Print("")
	os.Exit(0)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		usage()
	}

	// playlist
	if args[1] == "playlists" {
		playlist.Playlists()
		return
	}

	// queries
	if args[1] == "queries" {
		query.Queries()
		return
	}

	if len(args) < 3 {
		usage()
	}

	// inspect
	if (args[1] == "inspect") && (args[2] != "") {
		inspect.Inspect(args[2])
		return
	}

	// sync
	if (args[1] == "sync") && (args[2] != "") {
		sync.Sync(args[2])
		return
	}

	// show
	if (args[1] == "show") && (args[2] != "") {
		show.Show(args[2])
		return
	}

	// query
	if (args[1] == "query") && (args[2] != "") {
		query.Query(args[2])
		return
	}

	usage()
}
