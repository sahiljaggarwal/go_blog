package user_dto

import (
	"errors"
	"log"
	"regexp"
)

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginDTO) Validate() error {
	log.Println("dto running ", l)
	if l.Email == "" {
		return errors.New("email is required")
	} 

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(l.Email) {
		return errors.New("invalid email address")
	}

	if l.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
