package test

import (
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func BuildUserData() *entity.User {
	return &entity.User{
		Id:              1,
		Username:        "test",
		CryptedPassword: "ahfjsh123",
		Role:            "rw",
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}
}

func BuildConfigData() *entity.Config {
	return &entity.Config{
		Id:           1,
		Name:         "Robert",
		ConfigValues: "Hello",
		Version:      1,
		CreatedAt:    time.Now().UTC(),
		ActorId:      2,
	}
}
