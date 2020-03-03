package main

import (
	"../config"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Error reading config", err.Error())
	}
	log.Info("Obtaining of template from file")
	templateFile := getTemplate(cfg)
	log.Info("Template has obtained")

	log.Debug("Template length:", len(templateFile))
	offset := cfg.Builder.Offset

	s := templateFile[offset : offset+4]
	log.Info(net.IPv4(s[0], s[1], s[2], s[3]))

	ip := net.ParseIP("37.21.168.41")

	if ip == nil {
		log.Fatalln("Unable to parse IP")
	}
	s[0] = ip[15]
	s[1] = ip[14]
	s[2] = ip[13]
	s[3] = ip[12]
	//copy(s, ip[12:])
	log.Info(net.IPv4(s[0], s[1], s[2], s[3]))

}

func getTemplate(cfg *config.Config) []byte {
	fileName := cfg.Builder.TemplateFile
	templateFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	return templateFile
}
