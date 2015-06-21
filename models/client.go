package models

type Client struct {
	Id          string
	Secret      string
	RedirectUri string
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
