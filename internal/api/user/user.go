package user

import (
	"bitly-clone/configs/db"
	"bitly-clone/internal/helpers"
	"bitly-clone/internal/repository"
	"bitly-clone/models"
	"bitly-clone/models/requests"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {

	userReq := new(requests.UserRegister)

	_ = c.Bind(userReq)

	val := helpers.Validation(userReq, c)

	if val != nil {
		return nil
	}

	user := models.User{
		Username: userReq.Username,
		Password: userReq.Password,
	}

	checkUser := models.User{}
	repository.Get().User().FindByUsername(user.Username, &checkUser)

	if checkUser.ID > 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:        400,
			Description: "Girmiş olduğunuz kullanıcı adı sistemde kayıtlıdır.",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {

		return c.JSON(http.StatusBadRequest, models.Response{
			Code:        400,
			Description: "Şifre oluşturulurken hata ile karşılaşıldı.",
		})

	}

	user.Password = string(hashedPassword)
	user.Token = createToken()
	err = repository.Get().User().Create(&user)

	if err != nil {

		return c.JSON(http.StatusBadRequest, models.Response{
			Code:        400,
			Description: "Üyelik oluşturulamadı.",
		})

	}

	return c.JSON(http.StatusCreated, models.Response{
		Code:        201,
		Description: "Üyelik oluşturuldu.",
		Data:        user.Token,
	})

}

func GetToken(c echo.Context) error {
	user := new(requests.ShowToken)

	_ = c.Bind(user)

	val := helpers.Validation(user, c)

	if val != nil {
		return nil
	}

	checkUser := models.User{}
	repository.Get().User().FindByUsername(user.Username, &checkUser)

	if checkUser.ID == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code: 400,
			Data: "Kullanıcı bulunamadı.",
		})
	}
	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code: 400,
			Data: "Geçersiz şifre..",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Data: checkUser.Token,
	})
}

func createToken() string {
	str := ""
	check := true
	var user models.User
	for check {
		a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "a", "b", "c", "d", "e", "A", "B", "C", "D", "E", "X", "U", "L", "DF", "qX", "fgQo"}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		str = strings.Join(a, "")
		db.MyDB.Model(user).Where("token = ?", str).Scan(&user)
		if user.ID == 0 {
			check = false
		}
	}

	return str
}
