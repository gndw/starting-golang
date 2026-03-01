package inmemorydb

import (
	"context"
	"testing"

	"github.com/gndw/starting-golang/internals/models"
	"github.com/stretchr/testify/assert"
)

func TestImplementation_GetUserData(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name     string
		args     args
		want     models.User
		wantErr  bool
		errStr   string
	}{
		{
			name: "should return user when user ID exists",
			args: args{
				ctx:    context.Background(),
				userID: 100,
			},
			want: models.User{
				ID:       100,
				FullName: "John Doe",
			},
			wantErr: false,
		},
		{
			name: "should return error when user ID does not exist",
			args: args{
				ctx:    context.Background(),
				userID: 999,
			},
			want:    models.User{},
			wantErr: true,
			errStr:  "[GetUserData] user with ID 999 not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewRepository(context.Background())
			got, err := m.GetUserData(tt.args.ctx, tt.args.userID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errStr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestNewRepository(t *testing.T) {
	t.Run("should successfully create new repository", func(t *testing.T) {
		repo, err := NewRepository(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, repo)
		impl := repo.(*Implementation)
		assert.NotNil(t, impl.mu)
		assert.NotEmpty(t, impl.users)
	})
}
