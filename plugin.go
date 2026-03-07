package loglinter

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/gtimofej0303/loglinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

type Settings struct{
	ForbiddenPatterns []string `mapstructure:"forbidden-patterns"`
    ForbiddenWords    []string `mapstructure:"forbidden-words"`
}

type plugin struct{
	settings Settings
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	a := analyzer.NewAnalyzer(p.settings.ForbiddenPatterns, p.settings.ForbiddenWords)
	return []*analysis.Analyzer{a}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil{
		return nil, err
	}
	return &plugin{settings: s}, nil
}

func init() {
	register.Plugin("loglinter", New)
}

var _ register.LinterPlugin = new(plugin)
