package database

import "time"

type RecordSong struct {
	Path     string
	Meta     map[string]interface{}
	Datetime time.Time
}

type RecordQuery struct {
	Name  string
	Query string
}

type RecordQueryResult struct {
	Path  string
	Key   string
	Value string
}
