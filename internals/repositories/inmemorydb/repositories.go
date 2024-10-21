package inmemorydb

import (
	"context"

	"github.com/gndw/starting-golang/internals/models"
)

//go:generate mockery --name Repository
type Repository interface {
	GetUserData(ctx context.Context, userID int64) (response models.User, err error)
}
