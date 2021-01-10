package helpers

import (
	"bitly-clone/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"net/http"
)

func Validation(model interface{}, c echo.Context) error {
	validate := validator.New()
	err := validate.Struct(model)
	var errMessage = ""

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			errMessage += err.Field() + "(" + err.Type().String() + ") alanı gereklidir. "

		}

		c.JSON(http.StatusBadRequest, models.Validation{
			Code:    400,
			Message: errMessage,
		})

		return err
	}

	return nil
}
