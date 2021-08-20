package link

import (
	"bitly-clone/internal/helpers"
	"bitly-clone/internal/repository"
	"bitly-clone/models"
	"bitly-clone/models/requests"
	"bitly-clone/models/response"
	"net/http"

	"github.com/labstack/echo"
)

func Redirect(c echo.Context) error {
	url := c.Param("link")

	var link models.Link
	repository.Get().Link().FindByShortURL(nil, url, &link)

	if link.ID == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:        http.StatusBadRequest,
			Description: "Link bulunamadı..",
		})
	}

	c.Redirect(http.StatusMovedPermanently, link.Url)

	return nil
}

func Store(c echo.Context) error {

	linkReq := new(requests.LinkStore)

	_ = c.Bind(linkReq)

	user := helpers.GetUser(c)

	val := helpers.Validation(linkReq, c)

	if val != nil {
		return nil
	}

	var linkCheck models.Link
	repository.Get().Link().FindByURL(&user.ID, linkReq.Url, &linkCheck)

	if linkCheck.ID > 0 {
		return c.JSON(http.StatusOK, models.Response{
			Code:        http.StatusOK,
			Description: "Girmiş olduğunuz url sistemde zaten kayıtlı!",
			Data:        response.ShortUrlResponse{ShortUrl: linkCheck.ShortUrl},
		})
	}

	link := models.Link{
		UserID:   user.ID,
		Url:      linkReq.Url,
		ShortUrl: helpers.CreateLink(),
	}

	err := repository.Get().Link().Create(&link)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:        http.StatusBadRequest,
			Description: "Link oluşturulamadı.",
		})
	}

	return c.JSON(http.StatusCreated, models.Response{
		Code:        http.StatusCreated,
		Description: "Kısa link oluşturuldu.",
		Data:        link,
	})
}

func Delete(c echo.Context) error {

	linkReq := new(requests.LinkDelete)

	_ = c.Bind(linkReq)

	user := helpers.GetUser(c)

	val := helpers.Validation(linkReq, c)

	if val != nil {
		return nil
	}

	var linkCheck models.Link
	repository.Get().Link().FindByID(&user.ID, linkReq.ID, &linkCheck)

	if linkCheck.ID == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:        http.StatusBadRequest,
			Description: "Link bulunamadı..",
		})
	}

	err := repository.Get().Link().Delete(&linkCheck)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:        http.StatusBadRequest,
			Description: "Link silinemedi.",
		})
	}

	return c.JSON(http.StatusCreated, models.Response{
		Code:        http.StatusCreated,
		Description: "Link silindi.",
	})
}

func List(c echo.Context) error {
	user := helpers.GetUser(c)

	var links []models.Link

	query := repository.Get().Link().List(user.ID)

	res := helpers.Paginate(query, helpers.GetPage(c), &links)

	return c.JSON(http.StatusOK, res)
}
