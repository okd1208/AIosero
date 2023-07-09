package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetNextMovePosition(c echo.Context) error {
	// クライアントから送られてきたデータを受け取る
	userColor := c.QueryParam("userColor")
	positions := c.QueryParam("positions")

	fmt.Println(userColor, positions)
	return c.String(http.StatusOK, "userColor = "+userColor)
}
