package datasource

import (
	"github.com/villers/api/datamodel"
	"time"
)

// Movies is our imaginary data source.
var Users = map[int64]datamodel.User{
	1: {
		ID:        1,
		Login:     "mickael",
		Password:  "password",
		Email:     "mickael@google.com",
		Companies: []*datamodel.Company{},
		CreatedAt: time.Now(),
	},
	2: {
		ID:        2,
		Login:     "priscilla",
		Password:  "password",
		Email:     "priscilla@facebook.com",
		Companies: []*datamodel.Company{},
		CreatedAt: time.Now(),
	},
}
