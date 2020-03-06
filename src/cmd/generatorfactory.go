package main

import (
	"../config"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
)

type generatorFactory struct {
	template []byte
	offset   int
}

func (f *generatorFactory) createGenerator(insert []byte) Generator {
	return Generator{template: &f.template, insertOffset: f.offset, insert: insert}
}

func getGeneratorFactory(cfg *config.Config) generatorFactory {
	log.Info("Obtaining of template from file")
	templateFile := getFileTemplate(cfg)
	log.Info("Template has obtained")

	log.Debug("Template length:", len(templateFile))
	offset := cfg.Builder.Offset

	return generatorFactory{template: templateFile, offset: offset}
}

func getFileTemplate(cfg *config.Config) []byte {
	fileName := cfg.Builder.TemplateFile
	templateFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	return templateFile
}
