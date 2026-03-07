package loglinter

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/gtimofej0303/loglinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

type Settings struct {
	Patterns []string `mapstructure:"patterns"`
	Words    []string `mapstructure:"words"`
}

type plugin struct {
	settings Settings
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	a := analyzer.NewAnalyzer(p.settings.Patterns, p.settings.Words)
	return []*analysis.Analyzer{a}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

func New(raw any) (register.LinterPlugin, error) {
	m, ok := raw.(map[string]any)
	if !ok {
		return &plugin{settings: Settings{}}, nil
	}

	var s Settings

	if rawPatterns, ok := m["patterns"]; ok {
		if arr, ok := rawPatterns.([]any); ok {
			for _, v := range arr {
				if str, ok := v.(string); ok {
					s.Patterns = append(s.Patterns, str)
				}
			}
		}
	}

	if rawWords, ok := m["words"]; ok {
		if arr, ok := rawWords.([]any); ok {
			for _, v := range arr {
				if str, ok := v.(string); ok {
					s.Words = append(s.Words, str)
				}
			}
		}
	}

	return &plugin{settings: s}, nil
}

func init() {
	register.Plugin("loglinter", New)
}

var _ register.LinterPlugin = new(plugin)
