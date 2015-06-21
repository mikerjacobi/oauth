package controllers

import (
	"net/http"

	"github.com/RangelReale/osin"
	log "github.com/Sirupsen/logrus"
	"github.com/gocraft/web"
)

type AuthContext struct {
	osin.Server
	Name string
}

func (c *AuthContext) AuthorizeHandler(w web.ResponseWriter, r *web.Request) {
	log.Error(r.Request)

	resp := c.NewResponse()
	defer resp.Close()

	if ar := c.HandleAuthorizeRequest(resp, r.Request); ar != nil {

		// HANDLE LOGIN PAGE HERE

		ar.Authorized = true
		c.FinishAuthorizeRequest(resp, r.Request, ar)
	}
	osin.OutputJSON(resp, w, r.Request)
}

func (c *AuthContext) TokenHandler(w web.ResponseWriter, r *web.Request) {
	req, err := http.NewRequest(r.Method, r.RoutePath(), r.Body)
	if err != nil {
		log.Panic("failed to convert web.Request into http.Request")
	}

	resp := c.NewResponse()
	defer resp.Close()

	if ar := c.HandleAccessRequest(resp, req); ar != nil {
		ar.Authorized = true
		c.FinishAccessRequest(resp, req, ar)
	}
	osin.OutputJSON(resp, w, req)
}
