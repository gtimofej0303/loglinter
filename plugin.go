package loglinter

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/gtimofej0303/loglinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

type plugin struct{}

func (*plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}

func (*plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

func New(settings any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

func init() {
	register.Plugin("loglinter", New)
}

var _ register.LinterPlugin = new(plugin)
