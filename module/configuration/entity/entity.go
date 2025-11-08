package entity

import "database/sql"

type Entity struct {
	Configs ConfigEntity
	Users   UserEntity
}

func NewEntities(db *sql.DB) Entity {
	return Entity{
		Configs: ConfigEntity{DB: db},
		Users:   UserEntity{DB: db},
	}
}
