package inmemorydb

import (
	"context"

	"github.com/gndw/starting-golang/internals/models"
)

type Repository interface {
	GetUserData(ctx context.Context, userID int64) (response models.User, err error)
}
