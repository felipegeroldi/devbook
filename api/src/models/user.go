package models

import (
	"errors"
	"strings"
	"time"
)

// User struct represents an user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare will call the validate and format methods to prepare user data to storage
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Nickname == "" {
		return errors.New("nickname is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if step == "register" && user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)
}
