package inspect

import (
	"facette.io/natsort"
	"golang.org/x/exp/maps"
	"somnusalis.org/dj/filesystem"
	"somnusalis.org/dj/log"
)

func Inspect(path string) {
	log.Header("INSPECT", path)

	meta := filesystem.ReadMeta(path)
	keys := maps.Keys(meta)
	natsort.Sort(keys)
	for _, k := range keys {
		v := meta[k]
		log.Action(k, v)
	}

	log.Footer("DONE")
}
