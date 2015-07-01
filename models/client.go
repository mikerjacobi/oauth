package models

type Client struct {
	Id          string `db:"client_id"`
	Secret      string `db:"client_secret"`
	RedirectUri string `db:"redirect_uri"`
	UserData    interface{}
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
