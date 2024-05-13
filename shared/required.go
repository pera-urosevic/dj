package shared

import (
	"strings"

	"github.com/dhowden/tag"
)

func hasValidYear(meta map[string]interface{}) string {
	year, hasYear := meta["TYER"]
	if !hasYear {
		return "Missing Year"
	}
	yearString := year.(string)
	if len(yearString) != 4 {
		return "Invalid Year"
	}
	return ""
}

func hasValidCatalog(meta map[string]interface{}) string {
	for k, v := range meta {
		if strings.HasPrefix(k, "COMM") {
			description := v.(*tag.Comm).Description
			if strings.HasPrefix(description, "Catalog") {
				text := v.(*tag.Comm).Text
				if text == "Mobile" || text == "" {
					return ""
				} else {
					return "Invalid Catalog"
				}
			}
		}
	}
	return ""
}

func hasValidLyrics(meta map[string]interface{}) string {
	data, hasTag := meta["USLT"]
	if !hasTag {
		return "Missing Lyrics"
	}
	lyrics := data.(*tag.Comm).Text
	if len(lyrics) < 1 {
		return "Missing Lyrics"
	}
	return ""
}

func hasValidReplayGain(meta map[string]interface{}) string {
	for k, v := range meta {
		if strings.HasPrefix(k, "TXXX") {
			description := v.(*tag.Comm).Description
			if strings.HasPrefix(description, "replaygain_") {
				text := v.(*tag.Comm).Text
				if len(text) > 0 {
					return ""
				}
			}
		}
	}
	return "Missing ReplayGain"
}

func hasRequiredTags(meta map[string]interface{}) []string {
	warnings := []string{}

	requiredTags := map[string]string{
		"APIC": "Missing Cover",
		"TALB": "Missing Album",
		"TBPM": "Missing BPM",
		"TCON": "Missing Genre",
		"TIT2": "Missing Title",
		"TPE1": "Missing Artist",
		"TRCK": "Missing TrackNumber",
	}
	for tag, warning := range requiredTags {
		_, hasTag := meta[tag]
		if !hasTag {
			warnings = append(warnings, warning)
		}
	}

	warningValidYear := hasValidYear(meta)
	if warningValidYear != "" {
		warnings = append(warnings, warningValidYear)
	}

	warningRequiredLyrics := hasValidLyrics(meta)
	if warningRequiredLyrics != "" {
		warnings = append(warnings, warningRequiredLyrics)
	}

	warningReplayGain := hasValidReplayGain(meta)
	if warningReplayGain != "" {
		warnings = append(warnings, warningReplayGain)
	}

	warningCatalog := hasValidCatalog(meta)
	if warningCatalog != "" {
		warnings = append(warnings, warningCatalog)
	}

	return warnings
}
