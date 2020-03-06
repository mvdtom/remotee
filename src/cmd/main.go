package main

import (
	"../config"
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net"
	"net/http"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Error reading config", err.Error())
	}
	factory := getGeneratorFactory(cfg)

	launchServer(factory, cfg)
}

func launchServer(factory generatorFactory, cfg *config.Config) {
	e := echo.New()
	e.GET("/hostfile/:address", func(c echo.Context) error {
		pAddress := c.Param("address")
		e.Logger.Info("Parsing address:", pAddress)
		ip := net.ParseIP(pAddress)
		if ip == nil {
			//todo: right error handling
			e.Logger.Error("Unable to parse IP:", pAddress)
			return errors.New("Unable to parse address specified")
		}
		insert := []byte{ip[15], ip[14], ip[13], ip[12]}
		gen := factory.createGenerator(insert)
		c.Response().Header().Set("Content-Disposition", "Attachment;filename=tvnserver.exe")
		return c.Stream(http.StatusOK, echo.MIMEOctetStream, &gen)
	})
	//just for testing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})
	e.Logger.Fatal(e.Start(cfg.Server.Host + ":" + cfg.Server.Port))
}
