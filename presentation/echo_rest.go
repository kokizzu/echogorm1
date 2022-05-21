package presentation

import (
	"echogorm1/business"
	"github.com/labstack/echo/v4"
)

func ParseInput[T any](c echo.Context, in *T, cr *business.CommonRequest) error {
	err := c.Bind(&in)
	cookie, err := c.Cookie(`authToken`)
	if err == nil {
		cr.AuthToken = cookie.Value
	}
	return err
}

func RenderOutput[T any](c echo.Context, out *T, cr *business.CommonResponse) error {
	if cr.StatusCode == 0 {
		cr.StatusCode = 200
	}
	if cr.SetAuthToken != `` {
		// c.SetCookie(`authToken`, cr.SetAuthToken, 3600, ``, ``, false, false)
	}
	return c.JSON(cr.StatusCode, out)
}
