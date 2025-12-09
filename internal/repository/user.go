package repository

import (
	"database/sql"

	"glebosyatina/test_project/internal/domain"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(dbconn *sql.DB) *UserRepo {
	return &UserRepo{
		db: dbconn,
	}
}

func (ur *UserRepo) CreateUser(name string, surname string) (*domain.User, error) {
	var id uint64
	err := ur.db.QueryRow("INSERT INTO users (name, surname) VALUES ($1, $2) RETURNING id", name, surname).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		Id:      id,
		Name:    name,
		Surname: surname,
	}, nil
}
func (ur *UserRepo) GetUserById(id uint64) (*domain.User, error) {
	return nil, nil
}
func (ur *UserRepo) DeleteUserById(id uint64) error {
	return nil
}
