package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
)

//UsersRepository is an users repository
type UsersRepository struct {
	db *sql.DB
}

//NewUsersRepository creates an users repository
func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

// Create insert an new user in database
func (repository UsersRepository) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

// Search searchs user by name or nick
func (repository UsersRepository) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// SearchByEmail return an user searching by email
func (repository UsersRepository) SearchByEmail(email string) (models.User, error) {
	line, err := repository.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, nil
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// SearchByID return an user searching by id
func (repository UsersRepository) SearchByID(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?", ID,
	)
	if err != nil {
		return models.User{}, nil
	}
	defer lines.Close()

	var user models.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	if user.ID == 0 {
		return models.User{}, errors.New("User not found")
	}

	return user, nil
}

// Update updates an user by ID
func (repository UsersRepository) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Delete deletes and user by ID
func (repository UsersRepository) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from users where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
