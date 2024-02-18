package handler

import (
	"go-templ-practice/model"
	"go-templ-practice/view/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "asasd@gmail.com",
	}
	return render(c, user.Show(u))
}
