package lint

import (
	"fmt"
	"strings"

	"somnusalis.org/dj/shared"
)

func Lint(path string) {
	fmt.Println()
	fmt.Println("LINT", path)
	fmt.Println()

	files := listFiles(path)

	for _, file := range files {
		meta := shared.ReadMeta(file)
		warnings := shared.CheckMeta(meta)
		if len(warnings) > 0 {
			fmt.Println(file)
			fmt.Println(strings.Join(warnings, "\n"))
			fmt.Println()
		}
	}

	fmt.Println("DONE")
	fmt.Println()
}
