package business

import (
	"testing"

	"echogorm1/model"

	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"github.com/hexops/autogold"
)

func TestGuestLogin(t *testing.T) {
	g := Guest{}

	t.Run(`emailTooShort`, func(t *testing.T) {
		in := &Guest_LoginIn{}
		out := g.Login(in)

		want := autogold.Want(`emailTooShort1`, Guest_LoginOut{CommonResponse: CommonResponse{
			StatusCode: 400,
			ErrorMsg:   "email is too short",
		}})
		want.Equal(t, out)
	})

	t.Run(`passwordTooShort`, func(t *testing.T) {
		in := &Guest_LoginIn{
			Email: `test@gmail.com`,
		}
		out := g.Login(in)

		want := autogold.Want(`passwordTooShort1`, Guest_LoginOut{CommonResponse: CommonResponse{
			StatusCode: 400,
			ErrorMsg:   "password is too short",
		}})
		want.Equal(t, out)
	})

	t.Run(`userNotFoundMusterr`, func(t *testing.T) {
		g.GetUserByEmail = func(email string) (*model.User, error) {
			return nil, model.ErrUserNotFound
		}
		in := &Guest_LoginIn{
			Email:    `test@gmail.com`,
			Password: `test123`,
		}
		out := g.Login(in)

		want := autogold.Want(`userNotFoundMusterr1`, Guest_LoginOut{CommonResponse: CommonResponse{
			StatusCode: 400,
			ErrorMsg:   "user not found",
		}})
		want.Equal(t, out)
	})

	t.Run(`wrongPasswordMustError`, func(t *testing.T) {
		g.GetUserByEmail = func(email string) (*model.User, error) {
			user := &model.User{
				Email: email,
			}
			return user, nil
		}

		in := &Guest_LoginIn{
			Email:    `test@gmail.com`,
			Password: `test123`,
		}
		out := g.Login(in)

		want := autogold.Want(`wrongPasswordMustError1`, Guest_LoginOut{CommonResponse: CommonResponse{
			StatusCode: 400,
			ErrorMsg:   "wrong email or password",
		}})
		want.Equal(t, out)
	})

	t.Run(`correctPasswordMustReturnToken`, func(t *testing.T) {
		g.GetUserByEmail = func(email string) (*model.User, error) {
			user := &model.User{
				Email: email,
			}
			hash, _ := bcrypt.GenerateFromPassword([]byte(`test123`), bcrypt.DefaultCost)
			user.PasswordHash = string(hash)
			return user, nil
		}

		in := &Guest_LoginIn{
			Email:    `test`,
			Password: `test123`,
		}
		out := g.Login(in)

		assert.NotEmpty(t, out.SetAuthToken)
	})
}

func TestGuestRegister(t *testing.T) {
	g := Guest{}
	userMap := map[string]string{}
	g.GetUserByEmail = func(email string) (*model.User, error) {
		if pass, ok := userMap[email]; ok {
			hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
			return &model.User{
				Model:        gorm.Model{ID: 1},
				Email:        email,
				PasswordHash: string(hash),
			}, nil
		}
		return nil, model.ErrUserNotFound
	}
	g.InsertUser = func(email, password string) error {
		userMap[email] = password
		return nil
	}

	t.Run(`registerMustSucceed`, func(t *testing.T) {
		in := &Guest_RegisterIn{
			Email:    `test@gmail.com`,
			Password: `test123`,
		}
		out := g.Register(in)

		want := autogold.Want(`registerMustSucceed1`, Guest_RegisterOut{Success: true})
		want.Equal(t, out)
	})

	t.Run(`loginMustSucceed`, func(t *testing.T) {

		in := &Guest_LoginIn{
			Email:    `test@gmail.com`,
			Password: `test123`,
		}

		out := g.Login(in)

		assert.NotEmpty(t, out.SetAuthToken)
	})

	t.Run(`registerDuplicateMustFail`, func(t *testing.T) {
		in := &Guest_RegisterIn{
			Email:    `test@gmail.com`,
			Password: `test123`,
		}
		out := g.Register(in)

		want := autogold.Want(`registerDuplicateMustFail1`, Guest_RegisterOut{CommonResponse: CommonResponse{
			StatusCode: 400,
			ErrorMsg:   "user already exists",
		}})
		want.Equal(t, out)
	})

}
