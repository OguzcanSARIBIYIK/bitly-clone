package link

import (
	"bitly-clone/configs/db"
	"bitly-clone/internal/helpers"
	"bitly-clone/models"
	"bitly-clone/models/requests"
	"github.com/labstack/echo"
	"net/http"
)

func Store(c echo.Context) error {

	linkReq := new(requests.LinkStore)

	_ = c.Bind(linkReq)

	user := helpers.GetUser(c)

	val := helpers.Validation(linkReq, c)

	if val != nil {
		return nil
	}

	link := models.Link{
		UserID:   user.ID,
		Url:      linkReq.Url,
		ShortUrl: "http://localhost:8000/" + helpers.CreateLink(),
	}

	err := db.MyDB.Create(&link).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code: 400,
			Data: "Link oluşturulamadı.",
		})
	}

	return nil
}
