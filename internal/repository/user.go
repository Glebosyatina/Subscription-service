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
	var user domain.User
	err := ur.db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.Surname)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (ur *UserRepo) DeleteUserById(id uint64) error {
	if _, err := ur.db.Exec("DELETE FROM users WHERE id=$1 RETURNING id", id); err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) GetAllUsers() ([]*domain.User, error) {
	users := make([]*domain.User, 0)

	rows, err := ur.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.Id, &u.Name, &u.Surname)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}
