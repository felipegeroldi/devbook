package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// users represent an user repository
type users struct {
	db *sql.DB
}

// NewUserRepository create new user repositories
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) CreateUser(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users(Name, Nickname, Email, Password) VALUES(?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil

}

func (repository users) SearchUsers(nameOrNickname string) ([]models.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname)

	rows, err := repository.db.Query(
		"select Id, Name, Nickname, Email, CreatedAt FROM users WHERE Name LIKE ? or Nickname LIKE ?",
		nameOrNickname, nameOrNickname,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}

func (repository users) SearchUserById(userID uint64) (models.User, error) {
	row, err := repository.db.Query(
		"Select Id, Name, Nickname, Email, CreatedAt FROM users WHERE Id = ?",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) UpdateUser(userID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET Name = ?, Nickname = ?, Email = ? WHERE Id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nickname, user.Email, userID)
	if err != nil {
		return err
	}

	return nil
}
