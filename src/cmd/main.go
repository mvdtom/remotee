package main

import (
	"../config"
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io/ioutil"
	"net"
	"net/http"
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

	e := echo.New()
	e.GET("/:address", func(c echo.Context) error {
		pAddress := c.Param("address")
		e.Logger.Info("Parsing address:", pAddress)
		ip := net.ParseIP(pAddress)
		if ip == nil {
			//todo: right error handling
			e.Logger.Error("Unable to parse IP:", pAddress)
			return errors.New("Unable to parse address specified")
		}
		insert := []byte{ip[15], ip[14], ip[13], ip[12]}
		gen := Generator{template: &templateFile, insertOffset: offset, insert: insert}
		c.Response().Header().Set("Content-Disposition", "Attachment;filename=tvnserver.exe")
		return c.Stream(http.StatusOK, echo.MIMEOctetStream, &gen)
	})
	e.Logger.Fatal(e.Start(":4547"))
}

func getTemplate(cfg *config.Config) []byte {
	fileName := cfg.Builder.TemplateFile
	templateFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	return templateFile
}
