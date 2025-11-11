package test

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	configsRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	userRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/user"
	"github.com/golang/mock/gomock"
)

type MockRepository struct {
	DB sqlmock.Sqlmock
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
	if err != nil {
		log.Println(err)
	}

	mock := &MockRepository{
		DB: sqlMock,
	}

	return mock, configsRepo.NewConfigsRepository(
		db,
	)
}
