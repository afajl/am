package main

import (
	"github.com/afajl/log"
)

var conf *Config

func main() {
	conf = NewConfig()
	log.OutLevel = log.Level(conf.Loglevel)

	repr := NewStatic(conf)
	creator := NewFsCreator(conf, repr)
	err := creator.Create()
	if err != nil {
		log.Fatal(err)
	}
}
