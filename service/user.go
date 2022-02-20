package service

import (
	"github.com/wachayathorn/golang-graphql/config"
	"github.com/wachayathorn/golang-graphql/graph/model"
)

func CreateUser(data *model.User) (*model.User, error) {
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

func GetUserById(id int) (*model.User, error) {
	var user model.User
	err := config.SQLX.Get(&user, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(data *model.User) (*model.User, error) {
	tx := config.SQLX.MustBegin()

	err := tx.QueryRow(`UPDATE users SET firstname = $1, lastname = $2, username = $3 WHERE id = $4`, data.Firstname, data.Lastname, data.Username, data.ID).Scan()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return data, nil
}

func DeleteUser(id int) (*model.Response, error) {
	tx := config.SQLX.MustBegin()

	result, err := tx.Exec(`DELETE FROM users WHERE id = $1`, id)
	if row, _ := result.RowsAffected(); row == 0 || err != nil {
		tx.Rollback()
		return &model.Response{Success: false, Message: "User delete fail"}, err
	}
	tx.Commit()
	return &model.Response{Success: true, Message: "User deleted"}, nil
}
