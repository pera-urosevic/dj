package sync

import (
	"regexp"
	"unicode/utf8"
)

func truncate(s string, l int) string {
	if len(s) <= l {
		return s
	}
	var count int
	for i, char := range s {
		if count >= l {
			return string(s[:i])
		}
		count += utf8.RuneLen(char)
	}
	return s
}

func sanitize(filename string) string {
	reIllegal := regexp.MustCompile(`[\/\?<>\\:\*\|"]`)
	reControl := regexp.MustCompile(`[\x00-\x1f\x80-\x9f]`)
	reReserved := regexp.MustCompile(`^\.+$`)
	reWindowsReserve := regexp.MustCompile(`(?i)^(con|prn|aux|nul|com[0-9]|lpt[0-9])(\..*)?$`)
	reWindowsTrailing := regexp.MustCompile(`[\. ]+$`)
	t := filename
	t = reIllegal.ReplaceAllString(t, "")
	t = reControl.ReplaceAllString(t, "")
	t = reReserved.ReplaceAllString(t, "")
	t = reWindowsReserve.ReplaceAllString(t, "")
	t = reWindowsTrailing.ReplaceAllString(t, "")
	t = truncate(t, 240)
	return t
}
