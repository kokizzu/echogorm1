package business

import (
	"echogorm1/model"
	"github.com/kokizzu/gotro/L"
)

type Guest struct {
	GetUserByEmail func(email string) (*model.User, error)
	InsertUser     func(email, password string) error
}

type Guest_LoginIn struct {
	CommonRequest
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Guest_LoginOut struct {
	CommonResponse
	Token string `json:"token"`
}

const Guest_LoginRoute = `/guest/login`

func (g *Guest) Login(in *Guest_LoginIn) (out Guest_LoginOut) {
	if len(in.Email) < 3 {
		out.StatusCode = 400
		out.ErrorMsg = "email is too short"
		return
	}
	if len(in.Password) < 3 {
		out.StatusCode = 400
		out.ErrorMsg = "password is too short"
		return
	}
	user, err := g.GetUserByEmail(in.Email)
	if err != nil {
		out.StatusCode = 400
		out.ErrorMsg = "user not found"
		return
	}
	if !user.PasswordMatch(in.Password) {
		out.StatusCode = 400
		out.ErrorMsg = `wrong email or password`
		return
	}

	out.SetAuthToken = `123124`
	return
}

type Guest_RegisterIn struct {
	CommonRequest
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Guest_RegisterOut struct {
	CommonResponse
	Success bool `json:"success"`
}

const Guest_RegisterRoute = `/guest/register`

func (g *Guest) Register(in *Guest_RegisterIn) (out Guest_RegisterOut) {
	user, err := g.GetUserByEmail(in.Email)
	L.Describe(err)
	if err != nil && err.Error() != model.ErrUserNotFound.Error() {
		// TODO: check error and log
		out.StatusCode = 500
		out.ErrorMsg = err.Error() // TODO: censor
		return
	}
	if user != nil && user.ID != 0 {
		out.StatusCode = 400
		out.ErrorMsg = "user already exists"
		return
	}

	err = g.InsertUser(in.Email, in.Password)
	if err != nil {
		// TODO: check error and log
		out.StatusCode = 500
		out.ErrorMsg = err.Error() // TODO: censor
		return
	}

	out.Success = true
	return
}
