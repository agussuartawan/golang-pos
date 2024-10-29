package response

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AppendRolesResponse struct {
	UserId uint `json:"userId"`
	RoleIds []uint `json:"roleIds"`
}

type User struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone *string `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	Roles []Role `json:"roles,omitempty"`
}

type UserLoginResponse struct {
	ID uint
	Name string
	Password string
}

func (u *UserLoginResponse) ValidatePassword(plainPassword string) error {
	if u.Password == "" {
		return errors.New("invalid password")
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword)); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errors.New("invalid password")
	}

	return nil
}