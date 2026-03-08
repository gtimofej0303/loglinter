package rules

import (
	"strings"
	"unicode"
)

func AutoFixMessage(msg string) string {
	if msg == "" {
		return msg
	}

	runes := []rune(msg)
	if unicode.IsUpper(runes[0]) {
		runes[0] = unicode.ToLower(runes[0])
	}
	msg = string(runes)

	var b strings.Builder
	b.Grow(len(msg))

	for _, r := range msg {
		if unicode.Is(unicode.So, r) || unicode.Is(unicode.Sm, r) {
			continue
		}
		if strings.ContainsRune(forbidden, r) { //forbidden берём из specchars.go
			continue
		}

		b.WriteRune(r)
	}

	return b.String()
}
