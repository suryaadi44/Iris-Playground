package rest

import (
	"github.com/suryaadi44/iris-playground/app/api/rest/controller"
	repository "github.com/suryaadi44/iris-playground/app/repository/impl"
	usecase "github.com/suryaadi44/iris-playground/app/usecase/impl"
	"github.com/suryaadi44/iris-playground/utils/response"
	"github.com/suryaadi44/iris-playground/utils/validator"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func InitRoute(app *iris.Application, db *gorm.DB, conf *viper.Viper) {
	validator := validator.NewValidator()
	ur := repository.NewUserRepositoryImpl(db)
	us := usecase.NewUserServiceImpl(ur)
	uc := controller.NewUserController(us, validator)

	app.Get("/ping", Ping)

	app.Post("/signup", uc.SignUp)
	app.Post("/login", uc.LogIn)
}

func Ping(ctx iris.Context) {
	ctx.JSON(response.NewBaseResponse("pong", nil, nil))
}
