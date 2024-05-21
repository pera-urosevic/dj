package log

import (
	"fmt"

	"somnusalis.org/dj/database"
)

func Print(s string) {
	fmt.Println(s)
}

func Header(text ...string) {
	fmt.Printf("\n%s\n\n", text)
}

func Footer(text ...string) {
	fmt.Printf("\n%s\n\n", text)
}

func Action(action string, params ...interface{}) {
	fmt.Printf("[%s] %v\n", action, params)
}

func Warning(warning interface{}) {
	fmt.Printf("[Warning] [%v]\n", warning)
}

func Result(i int, result database.RecordSong) {
	fmt.Printf("%d. %s\n", i, result.Path)
}

func KeyValue(key string, value string) {
	fmt.Printf("%s: %v\n", key, value)
}
