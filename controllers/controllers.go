package controllers

import (
	"database/sql"

	"github.com/RangelReale/osin"
	"github.com/gocraft/web"
)

type AuthContext struct {
	osin.Server
	DB *sql.DB
}

func (c *AuthContext) HealthCheckHandler(w web.ResponseWriter, r *web.Request) {
	resp := c.NewResponse()
	defer resp.Close()
	resp.Output["hello"] = "world"
	osin.OutputJSON(resp, w, r.Request)
}

func (c *AuthContext) AuthorizeHandler(w web.ResponseWriter, r *web.Request) {
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
	resp := c.NewResponse()
	defer resp.Close()

	if ar := c.HandleAccessRequest(resp, r.Request); ar != nil {
		ar.Authorized = true
		c.FinishAccessRequest(resp, r.Request, ar)
	}
	osin.OutputJSON(resp, w, r.Request)
}
