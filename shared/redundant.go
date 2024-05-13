package shared

import (
	"strings"

	"github.com/dhowden/tag"
)

func hasComment(meta map[string]interface{}) string {
	for k, v := range meta {
		if strings.HasPrefix(k, "COMM") {
			description := v.(*tag.Comm).Description
			if description == "" {
				return "Invalid Comment"
			}
		}
	}
	return ""
}

func hasRedundantTags(meta map[string]interface{}) []string {
	warnings := []string{}

	warningComment := hasComment(meta)
	if warningComment != "" {
		warnings = append(warnings, warningComment)
	}

	return warnings
}
