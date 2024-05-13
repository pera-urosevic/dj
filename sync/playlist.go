package sync

import (
	"bytes"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func getPlaylist(playlistPath string) []string {
	raw, err := os.ReadFile(playlistPath)
	if err != nil {
		panic(err)
	}

	win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())
	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)
	decoded, err := io.ReadAll(unicodeReader)
	if err != nil {
		panic(err)
	}
	playlist := string(decoded)

	songs := []string{}
	lines := strings.Split(playlist, "\n")
	foundContent := false
	for _, line := range lines {
		if !foundContent {
			if strings.HasPrefix(line, "#-----CONTENT-----#") {
				foundContent = true
			}
		} else {
			if strings.HasPrefix(line, "#-") {
				break
			}
			if line == "" {
				break
			}
			trimmed := strings.TrimSpace(line)
			parts := strings.Split(trimmed, "|")
			filename := parts[0]
			songs = append(songs, filename)
		}
	}
	return songs
}
