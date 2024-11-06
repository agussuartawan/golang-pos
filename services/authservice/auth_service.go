package authservice

import (
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/payload"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/sessionrepository"
	"github.com/agussuartawan/golang-pos/repositories/userrepository"
	"os"
	"strconv"
	"time"
)

func Login(req request.LoginRequest, res *response.LoginResponse, session *payload.SessionCookie) error {
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

	// set response
	res.Name = user.Name
	for _, role := range user.Roles {
		res.Roles = append(res.Roles, role.Name)
		for _, permission := range role.Permissions {
			res.Permissions = append(res.Permissions, permission.Name)
		}
	}

	// generate jwt
	isSuperAdmin := helper.Contains(res.Roles, "super_admin")
	token, err := helper.CreateToken(sessionId, isSuperAdmin, expiredAt)
	if err != nil {
		return err
	}

	// set session
	session.SessionId = sessionId
	session.Token = token
	session.User = response.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return nil
}
