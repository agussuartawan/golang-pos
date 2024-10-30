package authservice

import (
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/sessionrepository"
	"github.com/agussuartawan/golang-pos/repositories/userrepository"
	"os"
	"strconv"
	"time"
)

func Login(req request.LoginRequest, res *response.LoginResponse) error {
	// prepare data from env
	JWTExpiration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		return err
	}

	// get user from DB
	var user models.User
	if err := userrepository.GetByEmail(&user, req.Email); err != nil {
		return err
	}

	// validate password
	if err := user.ValidatePassword(req.Password); err != nil {
		return err
	}

	// create session and clear previous session
	if err := sessionrepository.ClearSession(user.ID); err != nil {
		return err
	}
	expiredAt := time.Now().AddDate(0, 0, JWTExpiration)
	sessionId, err := sessionrepository.Create(user.ID, expiredAt, req.IpAddress)
	if err != nil {
		return err
	}

	// generate jwt
	token, err := helper.CreateToken(sessionId, expiredAt)
	if err != nil {
		return err
	}

	// mapping response
	res.Name = user.Name
	res.Token = token
	for _, role := range user.Roles {
		res.Roles = append(res.Roles, role.Name)
		for _, permission := range role.Permissions {
			res.Permissions = append(res.Permissions, permission.Name)
		}
	}

	return nil
}
