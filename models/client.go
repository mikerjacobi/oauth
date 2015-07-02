package models

import (
	"time"

	"github.com/RangelReale/osin"
)

type Client struct {
	Id          string
	Secret      string
	RedirectUri string
	UserData    interface{}
	Code        string
	ExpiresIn   int32
	State       string
	Created     time.Time
}

func (c Client) ToOsinAuthorizeData() osin.AuthorizeData {
	oad := osin.AuthorizeData{
		Client:      c,
		Code:        c.Code,
		Scope:       "",
		RedirectUri: c.RedirectUri,
		State:       c.State,
		CreatedAt:   c.Created,
		ExpiresIn:   c.ExpiresIn,
		UserData:    nil,
	}
	return oad
}

func (c Client) GetId() string {
	return c.Id
}

func (c Client) GetSecret() string {
	return c.Secret
}

func (c Client) GetRedirectUri() string {
	return c.RedirectUri
}

func (c Client) GetUserData() interface{} {
	return c.UserData
}
