package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/gtimofej0303/loglinter/pkg/analyzer/rules"
	
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type runner struct {
    extraPatterns []string
    extraWords    []string
}

func NewAnalyzer(extraPatterns []string, extraWords []string) *analysis.Analyzer {
    r := &runner{
        extraPatterns: extraPatterns,
        extraWords:    extraWords,
    }

    return &analysis.Analyzer{
        Name:     "loglinter",
        Doc:      "Checks that log messages follow formatting rules",
        Run:      r.run,
        Requires: []*analysis.Analyzer{inspect.Analyzer},
    }
}

var Analyzer = NewAnalyzer(nil, nil)

var loggerPkgs = map[string] bool{
	"slog": true,
	"zap": 	true,
	"log": 	true,
}

func (r *runner) run(pass *analysis.Pass) (interface{}, error){
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
        call, ok := n.(*ast.CallExpr)
        if !ok {
            return
        }

		msg, pos, ok := extractLogMessage(call)
		if !ok {
            return
        }

		rules.CheckLowercase(pass, msg, pos)
        rules.CheckEnglish(pass, msg, pos)
        rules.CheckSpecialChars(pass, msg, pos)
        rules.CheckSensitive(pass, msg, pos)

		if len(r.extraPatterns) > 0 || len(r.extraWords) > 0 {
            rules.CheckCustom(pass, msg, pos, r.extraPatterns, r.extraWords)
        }
	})

	return nil, nil
}

func extractLogMessage(call *ast.CallExpr) (string, token.Pos, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
    if !ok {
        return "", 0, false
    }

	methodName := sel.Sel.Name
    if !isLogMethod(methodName) {
        return "", 0, false
    }

	ident, ok := sel.X.(*ast.Ident)
    if !ok {
        return "", 0, false
    }
    if !loggerPkgs[ident.Name] && ident.Name != "logger" {
        return "", 0, false
    }

	if len(call.Args) == 0 {
        return "", 0, false
    }
	lit, ok := call.Args[0].(*ast.BasicLit)
    if !ok || lit.Kind != token.STRING {
        return "", 0, false
    }

	msg := strings.Trim(lit.Value, `"`)
    return msg, lit.Pos(), true
}

func isLogMethod(name string) bool{
	switch name {
    case "Info", "Error", "Warn", "Debug", "Fatal", "Panic", "Infof", "Errorf", "Warnf", "Debugf":
        return true
    }
    return false
}