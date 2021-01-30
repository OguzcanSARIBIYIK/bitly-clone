package link

import (
	"bitly-clone/configs/db"
	"bitly-clone/internal/helpers"
	"bitly-clone/models"
	"bitly-clone/models/requests"
	"bitly-clone/models/response"
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

	var linkCheck models.Link
	db.MyDB.Model(&models.Link{}).
		Where("user_id = ?", user.ID).
		Where("url = ?", linkReq.Url).
		First(&linkCheck)

	if linkCheck.ID > 0 {
		return c.JSON(http.StatusOK, models.Response{
			Code:        200,
			Description: "Girmiş olduğunuz url sistemde zaten kayıtlı!",
			Data:        response.ShortUrlResponse{ShortUrl: linkCheck.ShortUrl},
		})
	}

	link := models.Link{
		UserID:   user.ID,
		Url:      linkReq.Url,
		ShortUrl: helpers.CreateLink(),
	}

	err := db.MyDB.Create(&link).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:        400,
			Description: "Link oluşturulamadı.",
		})
	}

	return c.JSON(http.StatusCreated, models.Response{
		Code:        201,
		Description: "Kısa link oluşturuldu.",
		Data:        link,
	})
}

func List(c echo.Context) error {

	user := helpers.GetUser(c)

	var links []models.Link

	db.MyDB.Model(&models.Link{}).Where("user_id = ?", user.ID).Find(&links)

	return c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Data: links,
	})
}
