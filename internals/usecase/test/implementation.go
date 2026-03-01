package test

import (
	"context"

	"github.com/gndw/starting-golang/internals/models"
	"github.com/gndw/starting-golang/internals/repositories/inmemorydb"
)

type Implementation struct {
	inMemoryDbRepository inmemorydb.Repository
}

func NewUsecase(ctx context.Context, inMemoryDbRepository inmemorydb.Repository) (Usecase, error) {
	h := &Implementation{
		inMemoryDbRepository: inMemoryDbRepository,
	}
	return h, nil
}

func (m *Implementation) Test(ctx context.Context, request models.TestRequest) (response models.TestResponse, err error) {

	repositoryResponse, err := m.inMemoryDbRepository.GetUserData(ctx, request.UserID)
	if err != nil {
		return response, err
	}

	response.UserID = repositoryResponse.ID
	response.FullName = repositoryResponse.FullName
	return response, nil
}
