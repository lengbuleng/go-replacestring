package config

import (
	"github.com/spf13/viper"
)

type Parser struct {
	SrcPath string
	Ext     []string

	Keyword     []string
	Replaceword []string
}

func InitParser(cfg *viper.Viper) *Parser {

	return &Parser{
		SrcPath:     cfg.GetString("srcpath"),
		Ext:         cfg.GetStringSlice("ext"),
		Keyword:     cfg.GetStringSlice("keyword"),
		Replaceword: cfg.GetStringSlice("Replaceword"),
	}
}

var ParserConfig = new(Parser)
