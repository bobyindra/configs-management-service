package user

import (
	"database/sql"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

var (
	userColumn = []string{
		"id",
		"username",
		"crypted_password",
		"role",
		"created_at",
		"updated_at",
	}
)

type userRepo struct{ db *sql.DB }

func NewUserRepository(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

type userRecord struct {
	Id              uint      `db:"id"`
	Username        string    `db:"username"`
	CryptedPassword string    `db:"crypted_password"`
	Role            string    `db:"role"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

func (ur userRecord) ToEntity() *entity.User {
	uRes := &entity.User{
		Id:              ur.Id,
		Username:        ur.Username,
		CryptedPassword: ur.CryptedPassword,
		Role:            ur.Role,
		CreatedAt:       ur.CreatedAt,
		UpdatedAt:       ur.UpdatedAt,
	}
	return uRes
}
