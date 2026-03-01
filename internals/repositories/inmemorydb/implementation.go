package inmemorydb

import (
	"context"
	"fmt"
	"sync"

	"github.com/gndw/starting-golang/internals/models"
)

type Implementation struct {
	mu    *sync.Mutex
	users []models.User
}

func NewRepository(ctx context.Context) (Repository, error) {
	h := &Implementation{
		mu: &sync.Mutex{},
		users: []models.User{
			{
				ID:       100,
				FullName: "John Doe",
			},
		},
	}
	return h, nil
}

func (m *Implementation) GetUserData(ctx context.Context, userID int64) (response models.User, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, user := range m.users {
		if user.ID == userID {
			return user, nil
		}
	}
	return response, fmt.Errorf("[GetUserData] user with ID %v not found", userID)
}
