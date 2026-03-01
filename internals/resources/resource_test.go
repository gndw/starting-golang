package resources_test

import (
	"context"
	"os"
	"testing"

	"github.com/gndw/starting-golang/internals/resources"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	// Create a temporary .env file for the test
	err := os.WriteFile(".env", []byte("PORT=8080"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(".env")

	tests := []struct {
		name    string // description of this test case
		wantErr bool
	}{
		{
			name:    "should successfully init resources when all dependencies are available",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := resources.Init(context.Background())
			if tt.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
				assert.NotNil(t, got.HttpServerService)
			}
		})
	}
}
