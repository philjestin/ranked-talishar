package data

import (
	"context"
	"time"

	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
)

type Permissions []string

func (p Permissions) Include(code string) bool {
	for i := range p {
		if code == p[i] {
			return true
		}
	}
	return false
}

type PermissionModel struct {
	DB *db.Queries
}

func (m PermissionModel) GetAllForUser(userID uuid.UUID) (Permissions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.GetAllPermissionsForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (m PermissionModel) AddForUser(userID uuid.UUID, codes string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := &db.AddPermissionForUserParams{
		UserID: userID,
		Code:   codes,
	}

	err := m.DB.AddPermissionForUser(ctx, *args)
	return err
}
