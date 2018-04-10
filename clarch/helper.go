package clarch

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func mkdir(filePath string) error {
	return os.MkdirAll(filePath, os.ModePerm)
}

func openOrCreate(filePath string) (*os.File, error) {
	paths := strings.Split(filePath, "/")
	fileName := paths[len(paths)-1]
	dirs := strings.Join(paths[0 : len(paths)-1])
	if err := os.MkdirAll(dirs); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return os.OpenFile(fmt.Sprintf(filePath), os.O_RDWR|os.O_CREATE, 0666)
}

func camelCase(s string) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s))

	var prev rune
	for _, curr := range s {
		if !isDelimiter(curr) {
			if isDelimiter(prev) || (prev == 0) {
				buffer = append(buffer, toUpper(curr))
			} else {
				buffer = append(buffer, toLower(curr))
			}
		}
		prev = curr
	}

	return string(buffer)
}

func isDelimiter(ch rune) bool {
	return ch == '-' || ch == '_' || isSpace(ch)
}

func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
