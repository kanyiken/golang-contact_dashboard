package models

import "errors"

var ErrNoRecord = errors.New("models: no such record found")
var ErrInvalidCredentials = errors.New("models: invalid credentials")
var ErrDuplicateNumber = errors.New("models: number exists")

type (
	User struct {
		ID             int
		Name, Mobile   string
		HashedPassword []byte
	}

	Contact struct {
		ID           int
		Name, Mobile string
	}

	Groups struct {
		ID   int
		Name string
	}

	GroupedContacts struct {
		ID        int
		ContactID int
		GroupID   int
	}
)
