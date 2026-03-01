package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		m       *TestRequest
		wantErr bool
		errStr  string
	}{
		{
			name: "should return error when user ID is <= 0",
			m: &TestRequest{
				UserID: 0,
			},
			wantErr: true,
			errStr:  "[TestRequest] user_id cannot be <= 0",
		},
		{
			name: "should return error when user ID is negative",
			m: &TestRequest{
				UserID: -1,
			},
			wantErr: true,
			errStr:  "[TestRequest] user_id cannot be <= 0",
		},
		{
			name: "should return nil when user ID is > 0",
			m: &TestRequest{
				UserID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.m.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errStr, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
