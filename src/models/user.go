package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User models an user using application
type User struct {
	ID        uint64    `json:"id,omitempty"` // omitempty vai omitir o campo id se estiver vazio, ao invés de retornar com valor zero
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare validates and format received user
func (user *User) Prepare(action string) error {
	if err := user.validate(action); err != nil {
		return err
	}

	if err := user.format(action); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(action string) error {
	if user.Name == "" {
		return errors.New("Invalid blank name")
	}

	if user.Nick == "" {
		return errors.New("Invalid blank nick")
	}

	if user.Email == "" {
		return errors.New("Invalid blank email")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Invalid email format")
	}

	if action == "create" && user.Password == "" {
		return errors.New("Invalid blank password")
	}

	return nil
}

func (user *User) format(action string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if action == "create" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return nil
}
