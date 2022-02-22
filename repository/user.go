package repository

import (
	"github.com/wachayathorn/golang-graphql/config"
	"github.com/wachayathorn/golang-graphql/graph/model"
)

type UserRepositoryInterface interface {
	CreateUser(input *model.User) (*model.User, error)
	GetUsers() ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
	UpdateUser(input *model.User) (*model.User, error)
	DeleteUser(id int) (*model.Response, error)
}

type UserRepository struct{}

func (r UserRepository) CreateUser(input *model.User) (*model.User, error) {
	tx := config.SQLX.MustBegin()
	var user model.User
	err := tx.QueryRow(`INSERT INTO users(firstname, lastname, username) VALUES($1, $2, $3) RETURNING id`, input.Firstname, input.Lastname, input.Username).Scan(&user.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, err
}

func (r UserRepository) GetUsers() ([]*model.User, error) {
	var users []*model.User
	userList, err := config.SQLX.Query(`SELECT * FROM users`)

	if err != nil {
		return nil, err
	}

	for userList.Next() {
		var user model.User
		err = userList.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r UserRepository) GetUserById(id int) (*model.User, error) {
	var user model.User
	err := config.SQLX.Get(&user, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) UpdateUser(input *model.User) (*model.User, error) {
	tx := config.SQLX.MustBegin()

	var user model.User
	err := tx.QueryRow(`UPDATE users SET firstname = $1, lastname = $2, username = $3 WHERE id = $4`, input.Firstname, input.Lastname, input.Username, input.ID).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

func (r UserRepository) DeleteUser(id int) (*model.Response, error) {
	tx := config.SQLX.MustBegin()

	result, err := tx.Exec(`DELETE FROM users WHERE id = $1`, id)
	if row, _ := result.RowsAffected(); row == 0 || err != nil {
		tx.Rollback()
		return &model.Response{Success: false, Message: "User delete fail"}, err
	}
	tx.Commit()

	return &model.Response{Success: true, Message: "User deleted"}, nil
}
