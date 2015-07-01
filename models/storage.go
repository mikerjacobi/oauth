package models

import (
	"database/sql"

	"github.com/RangelReale/osin"
	"github.com/Sirupsen/logrus"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

// Clone the storage if needed. For example, using mgo, you can clone the session with session.Clone
// to avoid concurrent access problems.
// This is to avoid cloning the connection at each method access.
// Can return itself if not a problem.
func (s *Storage) Clone() osin.Storage {
	return s
}

// Close the resources the Storage potentially holds (using Clone for example)
func (s *Storage) Close() {}

// GetClient loads the client by id (client_id)
func (s *Storage) GetClient(id string) (osin.Client, error) {
	var client Client
	logrus.Info("start get client")
	err := s.DB.QueryRow("SELECT * FROM clients WHERE client_id = $1", id).Scan(
		&client.Id, &client.Secret, &client.RedirectUri)

	if err != nil {
		logrus.Infof("error get client: %s", err.Error())
		return client, err
	}
	logrus.Info("finish get client")
	return client, nil
}

// SaveAuthorize saves authorize data.
func (s *Storage) SaveAuthorize(*osin.AuthorizeData) error {
	return nil
}

// LoadAuthorize looks up AuthorizeData by a code.
// Client information MUST be loaded together.
// Optionally can return error if expired.
func (s *Storage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	data := osin.AuthorizeData{}
	return &data, nil
}

// RemoveAuthorize revokes or deletes the authorization code.
func (s *Storage) RemoveAuthorize(code string) error {
	return nil
}

// SaveAccess writes AccessData.
// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
func (s *Storage) SaveAccess(*osin.AccessData) error {
	return nil
}

// LoadAccess retrieves access data by token. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *Storage) LoadAccess(token string) (*osin.AccessData, error) {
	data := osin.AccessData{}
	return &data, nil
}

// RemoveAccess revokes or deletes an AccessData.
func (s *Storage) RemoveAccess(token string) error {
	return nil
}

// LoadRefresh retrieves refresh AccessData. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *Storage) LoadRefresh(token string) (*osin.AccessData, error) {
	data := osin.AccessData{}
	return &data, nil
}

// RemoveRefresh revokes or deletes refresh AccessData.
func (s *Storage) RemoveRefresh(token string) error {
	return nil
}
