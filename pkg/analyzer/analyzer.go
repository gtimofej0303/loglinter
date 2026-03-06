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

var Analyzer = &analysis.Analyzer{
	Name: 		"loglinter",
	Doc: 		"Checks that log messages follow formatting rules",
	Run: 		run, 
	Requires: 	[]*analysis.Analyzer{inspect.Analyzer},
}

var loggerPkgs = map[string] bool{
	"slog": true,
	"zap": 	true,
	"log": 	true,
}

func run(pass *analysis.Pass) (interface{}, error){
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