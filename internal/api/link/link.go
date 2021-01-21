package link

import (
	"bitly-clone/internal/helpers"
	"bitly-clone/models/requests"
	"github.com/labstack/echo"
)

func Store(c echo.Context) error {

	linkReq := new(requests.LinkStore)

	_ = c.Bind(linkReq)

	val := helpers.Validation(linkReq, c)

	if val != nil {
		return nil
	}

	return nil
}
