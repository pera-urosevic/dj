package info

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"somnusalis.org/dj/shared"
)

func Info(path string) {
	fmt.Println()
	fmt.Println("INFO", path)
	fmt.Println()

	meta := shared.ReadMeta(path)

	keys := []string{}
	for k := range meta {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		json, err := json.Marshal(meta[k])
		if err != nil {
			panic(err)
		}
		v := string(json)
		if len(v) > 100 {
			v = v[:100] + "…"
		}
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Println()

	warnings := shared.CheckMeta(meta)
	if len(warnings) > 0 {
		fmt.Println("WARNINGS")
		fmt.Println(strings.Join(warnings, "\n"))
		fmt.Println()
	}

	fmt.Println("DONE")
	fmt.Println()
}
