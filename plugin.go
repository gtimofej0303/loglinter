package loglinter

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/gtimofej0303/loglinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

type Settings struct {
	Enabled bool `mapstructure:"enabled"`

	EnableLowercase *bool `mapstructure:"enable_lowercase"`
	EnableEnglish   *bool `mapstructure:"enable_english"`
	EnableSpecChars *bool `mapstructure:"enable_specchars"`
	EnableSensitive *bool `mapstructure:"enable_sensitive"`
	EnableCustom    *bool `mapstructure:"enable_custom"`

	Patterns []string `mapstructure:"patterns"`
	Words    []string `mapstructure:"words"`
}

type plugin struct {
	settings Settings
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	if !p.settings.Enabled {
		return nil, nil
	}

	cfg := analyzer.Config{
		EnableLowercase: boolOrDefault(p.settings.EnableLowercase, true),
		EnableEnglish:   boolOrDefault(p.settings.EnableEnglish, true),
		EnableSpecChars: boolOrDefault(p.settings.EnableSpecChars, true),
		EnableSensitive: boolOrDefault(p.settings.EnableSensitive, true),
		EnableCustom:    boolOrDefault(p.settings.EnableCustom, false),

		ExtraPatterns: p.settings.Patterns,
		ExtraWords:    p.settings.Words,
	}

	a := analyzer.NewAnalyzer(cfg)
	return []*analysis.Analyzer{a}, nil
}

func boolOrDefault(v *bool, def bool) bool {
	if v == nil {
		return def
	}
	return *v
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

	if v, ok := m["enabled"].(bool); ok {
		s.Enabled = v
	} else {
		s.Enabled = true //default
	}

	parseBoolPtr := func(key string) *bool {
		if v, ok := m[key]; ok {
			if b, ok := v.(bool); ok {
				return &b
			}
		}
		return nil
	}

	s.EnableLowercase = parseBoolPtr("enable_lowercase")
	s.EnableEnglish = parseBoolPtr("enable_english")
	s.EnableSpecChars = parseBoolPtr("enable_specchars")
	s.EnableSensitive = parseBoolPtr("enable_sensitive")
	s.EnableCustom = parseBoolPtr("enable_custom")

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
