package main

import (
	"bufio"
	"fmt"
	"os"

	"somnusalis.org/dj/info"
	"somnusalis.org/dj/lint"
	"somnusalis.org/dj/sync"
)

func usage() {
	fmt.Println("usage:")
	fmt.Println("dj lint <folder>")
	fmt.Println("dj sync <playlist> <folder>")
	os.Exit(0)
}

func wait() {
	fmt.Print("Press enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	args := os.Args

	if len(args) < 3 {
		usage()
	}

	// info
	if (args[1] == "info") && (args[2] != "") {
		info.Info(args[2])
		wait()
		return
	}

	// lint
	if (args[1] == "lint") && (args[2] != "") {
		lint.Lint(args[2])
		wait()
		return
	}

	if len(args) < 4 {
		usage()
	}

	// sync
	if (args[1] == "sync") && (args[2] != "") && (args[3] != "") {
		sync.Sync(args[2], args[3])
		wait()
		return
	}

	usage()
}
