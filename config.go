package main

import (
	"flag"
)

type Config struct {
	SourceDir string
	TargetDir string
	Loglevel  int
	TemplDir      string
	PicAlternates map[string]int
}

func NewConfig() *Config {
	config := new(Config)
	flag.StringVar(&config.SourceDir, "source", "", "source dir")
	flag.StringVar(&config.TargetDir, "target", "", "target dir")
	flag.StringVar(&config.TemplDir, "templ", "templ", "template dir")
	flag.IntVar(&config.Loglevel, "level", 3, "log level [0=fatal...5=trace]")

	config.PicAlternates = map[string]int{
		"small":  300,
		"medium": 600,
		"large":  1200,
		"xlarge": 2000}

	flag.Parse()
	return config
}
