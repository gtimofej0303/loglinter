package rules

import (
	"go/token"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func CheckLowercase(pass *analysis.Pass, msg string, pos token.Pos) {
	if len(msg) == 0{
		return
	}

	runes := []rune(msg)
	if unicode.IsUpper(runes[0]){
		pass.Reportf(pos, "log message must start with a lowercase letter: %q", msg)
	}
}