package rules

import (
    "go/token"
    "unicode"

    "golang.org/x/tools/go/analysis"
)

func CheckEnglish(pass *analysis.Pass, msg string, pos token.Pos) {
    for _, r := range msg {
        if r > 127 && unicode.IsLetter(r) {
            pass.Reportf(pos, "log message must contain only English letters: %q", msg)
            return
        }
    }
}