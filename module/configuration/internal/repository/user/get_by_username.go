package user

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var user userRecord

	query := fmt.Sprintf("SELECT %s from users where username = $1", strings.Join(userColumn, ", "))
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.Id, &user.Username, &user.CryptedPassword, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound(username)
		}
		return nil, err
	}

	return util.GeneralNullable(*user.ToEntity()), nil
}
