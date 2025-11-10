package user

import "database/sql"

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
