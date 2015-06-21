package main

import (
	"fmt"
	"net/http"

	"github.com/RangelReale/osin"
	log "github.com/Sirupsen/logrus"
	"github.com/gocraft/web"
	"github.com/mikerjacobi/oauth/controllers"
	"github.com/mikerjacobi/oauth/models"
)

func main() {
	// TestStorage implements the "osin.Storage" interface
	s := models.NewStorage()
	osinConfig := osin.NewServerConfig()
	osinConfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	server := osin.NewServer(osinConfig, s)

	context := controllers.AuthContext{*server, "my auth context"}
	router := web.New(context).
		Get("/authorize", context.AuthorizeHandler).
		Get("/token", context.TokenHandler)

	port := "14000"
	log.Infof("Listening on port: %s", port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), router)
}
