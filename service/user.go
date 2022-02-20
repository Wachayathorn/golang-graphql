package service

import (
	"context"

	"github.com/wachayathorn/golang-graphql/config"
	"github.com/wachayathorn/golang-graphql/graph/model"
)

func CreateUser(ctx context.Context, data *model.User) (*model.User, error) {
	tx := config.SQLX.MustBegin()
	var id int64
	err := tx.QueryRow(`INSERT INTO users(firstname, lastname, username) VALUES($1, $2, $3) RETURNING id`, data.Firstname, data.Lastname, data.Username).Scan(&id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	user := &model.User{ID: int(id), Firstname: data.Firstname, Lastname: data.Lastname, Username: data.Username}
	return user, nil
}

func GetUsers() ([]*model.User, error) {
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
