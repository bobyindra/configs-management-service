package test

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	configsRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	userRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/user"
	"github.com/go-redis/redismock/v9"
	"github.com/golang/mock/gomock"
)

type MockRepository struct {
	DB    sqlmock.Sqlmock
	Cache redismock.ClientMock
}

func NewMockUserRepository(ctrl *gomock.Controller) (*MockRepository, repository.UserRepository) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
	}

	mock := &MockRepository{
		DB: sqlMock,
	}

	return mock, userRepo.NewUserRepository(
		db,
	)
}

func NewMockConfigsRepository(ctrl *gomock.Controller) (*MockRepository, repository.ConfigsManagementRepository) {
	db, sqlMock, err := sqlmock.New()
	cache, cacheMock := redismock.NewClientMock()
	if err != nil {
		log.Println(err)
	}

	mock := &MockRepository{
		DB:    sqlMock,
		Cache: cacheMock,
	}

	return mock, configsRepo.NewConfigsRepository(
		db,
		cache,
	)
}
