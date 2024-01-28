package repository

import (
	"database/sql"
)

type UserRepository struct {
	db  *sql.DB
	err error
}

func NewUserRepository() {
	return UserRepository
}
