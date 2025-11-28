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

type MockDBRepository struct {
	DB sqlmock.Sqlmock
}

type MockCacheRepository struct {
	Cache redismock.ClientMock
}

func NewMockUserRepository(ctrl *gomock.Controller) (*MockDBRepository, repository.UserRepository) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
	}

	mock := &MockDBRepository{
		DB: sqlMock,
	}

	return mock, userRepo.NewUserRepository(
		db,
	)
}

func NewMockConfigsDBRepository(ctrl *gomock.Controller) (*MockDBRepository, repository.ConfigsManagementDBRepository) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
	}

	mock := &MockDBRepository{
		DB: sqlMock,
	}

	return mock, configsRepo.NewConfigsDBRepository(
		db,
	)
}

func NewMockConfigsCacheRepository(ctrl *gomock.Controller) (*MockCacheRepository, repository.ConfigsManagementCacheRepository) {
	cache, cacheMock := redismock.NewClientMock()

	mock := &MockCacheRepository{
		Cache: cacheMock,
	}

	return mock, configsRepo.NewConfigsCacheRepository(
		cache,
	)
}
