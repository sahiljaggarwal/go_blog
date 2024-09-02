package user_dto

import (
	"errors"
	"regexp"
)

type SignUpDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (s *SignUpDTO) Validate() error {
	if s.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(s.Email) {
		return errors.New("invalid email address")
	}

	if s.Password == ""{
		return errors.New("password is required")
	}

	if s.Username == "" {
		return errors.New("username is required")
	}

	return nil



}