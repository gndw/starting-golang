package test

import (
	"context"

	"github.com/gndw/starting-golang/internals/models"
)

type Usecase interface {
	Test(ctx context.Context, request models.TestRequest) (response models.TestResponse, err error)
}
