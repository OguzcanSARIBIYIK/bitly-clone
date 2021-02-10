package helpers

import (
	"bitly-clone/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type PaginateSt struct {
	CurrentPage int         `json:"current_page"`
	NextPage    int         `json:"next_page"`
	Data        interface{} `json:"data"`
	LastPage    int         `json:"last_page"`
	TotalRecord int64       `json:"total_record"`
}

func Paginate(db *gorm.DB, page int, model interface{}) models.Response {

	var out PaginateSt

	db.Debug().Count(&out.TotalRecord)

	if out.TotalRecord != 0 {
		out.LastPage = int(math.Ceil(float64(out.TotalRecord) / 10))
	}

	if page == 0 {
		page = 1
	}

	if (page + 1) > out.LastPage {
		out.CurrentPage = out.LastPage
		out.NextPage = out.LastPage
	} else {
		out.CurrentPage = page
		out.NextPage = page + 1
	}

	offset := 0
	limit := 10 * out.CurrentPage
	if page != 0 {
		offset = (out.CurrentPage - 1) * 10
	}

	db.Limit(limit).
		Offset(offset).
		Find(model)

	out.Data = model

	res := models.Response{
		Code: 200,
		Data: out,
	}

	return res
}

func GetPage(c echo.Context) int {
	page, _ := strconv.Atoi(c.QueryParam("page"))

	return page
}
