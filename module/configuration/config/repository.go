package config

import (
	"database/sql"

	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
)

type repositoryList struct {
	configsManagement repository.ConfigsManagementRepository
}

func NewRepositoryList(db *sql.DB) repositoryList {
	return repository.NewConfigsManagementRepository(db)
}
