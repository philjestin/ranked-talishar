package data

import (
	"errors"

	dbCon "github.com/philjestin/ranked-talishar/db/sqlc"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Tokens      TokenModel
	Users       UserModel
	Permissions PermissionModel
	Heroes      HeroModel
}

func NewModels(db *dbCon.Queries) Models {
	return Models{
		Permissions: PermissionModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Users:       UserModel{DB: db},
		Heroes:      HeroModel{DB: db}, // Initialize the Heroes field
	}
}
