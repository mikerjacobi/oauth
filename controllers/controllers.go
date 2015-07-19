package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"github.com/RangelReale/osin"
	log "github.com/Sirupsen/logrus"
	"github.com/gocraft/web"
)

type AuthContext struct {
	osin.Server
	DB *sql.DB
}

type badCreds struct{ error }

func (c *AuthContext) HealthCheckHandler(w web.ResponseWriter, r *web.Request) {
	resp := c.NewResponse()
	defer resp.Close()
	resp.Output["hello"] = "world"
	osin.OutputJSON(resp, w, r.Request)
}

func (c *AuthContext) checkSession(session string) (bool, error) {
	return true, nil
}
func (c *AuthContext) checkCredentials(r *web.Request) (bool, error) {
	r.ParseForm()
	log.Infof(">>> %+v\n", r.Form)
	username := r.PostForm["username"][0]
	password := r.PostForm["password"][0]

	if username == "test" && password == "test" {
		return true, nil
	}

	return false, badCreds{errors.New("creds invalid")}
}

func (c *AuthContext) setSessionCookie(w web.ResponseWriter) {
}

func (c *AuthContext) checkLogin(w web.ResponseWriter, r *web.Request) bool {
	/*
		this function should only be called from AuthorizeHandler (for now)
		1) check for valid session
			a) if valid, return true
		2) elseif  request.Method == POST and user/pw valid, return true
		3) else return false
	*/
	log.Info(1)
	//2)
	if r.Method == "POST" {
		log.Info(2)
		isValidCreds, _ := c.checkCredentials(r)
		c.setSessionCookie(w)
		return isValidCreds
	}

	//1)
	sessionCookie, err := r.Cookie("session")
	if err != nil {
		log.Info(3)
		return false
	}

	isValidSession, err := c.checkSession(sessionCookie.String())
	if err != nil {
		log.Info(4)
		return false
	} else if isValidSession {
		log.Info(5)
		return true
	}

	//3)
	return false
}

func (c *AuthContext) LoginHandler(w web.ResponseWriter, r *web.Request) {
}

func (c *AuthContext) AuthorizeHandler(w web.ResponseWriter, r *web.Request) {
	log.Error("In Auth Handler")
	resp := c.NewResponse()
	defer resp.Close()

	log.Infof("request form: %+v\n", r.Request.Form)
	if ar := c.HandleAuthorizeRequest(resp, r.Request); ar != nil {

		// HANDLE LOGIN PAGE HERE
		log.Info("a1")
		if c.checkLogin(w, r) {
			log.Info("a2")
			ar.Authorized = true
			c.FinishAuthorizeRequest(resp, r.Request, ar)
		} else {
			log.Info("a3")
			w.Write([]byte("<html><body>"))

			w.Write([]byte(fmt.Sprintf("LOGIN %s (use test/test)<br/>", ar.Client.GetId())))
			w.Write([]byte(fmt.Sprintf("<form action=\"/authorize?response_type=%s&client_id=%s&state=%s&redirect_uri=%s\" method=\"POST\">",
				ar.Type, ar.Client.GetId(), ar.State, url.QueryEscape(ar.RedirectUri))))

			w.Write([]byte("Login: <input type=\"text\" name=\"username\" /><br/>"))
			w.Write([]byte("Password: <input type=\"password\" name=\"password\" /><br/>"))
			w.Write([]byte("<input type=\"submit\"/>"))

			w.Write([]byte("</form>"))

			w.Write([]byte("</body></html>"))
			return
		}
	}
	osin.OutputJSON(resp, w, r.Request)
}

func (c *AuthContext) TokenHandler(w web.ResponseWriter, r *web.Request) {
	log.Error("In Token Handler")
	resp := c.NewResponse()
	defer resp.Close()

	if ar := c.HandleAccessRequest(resp, r.Request); ar != nil {
		ar.Authorized = true
		c.FinishAccessRequest(resp, r.Request, ar)
	}
	osin.OutputJSON(resp, w, r.Request)
}
