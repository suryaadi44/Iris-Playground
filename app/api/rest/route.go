package rest

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"suryaadi44/iris-playground/app/api/rest/controller"
	repository "suryaadi44/iris-playground/app/repository/impl"
	service "suryaadi44/iris-playground/app/usecase/impl"
	"suryaadi44/iris-playground/utils/response"
	"suryaadi44/iris-playground/utils/validator"
)

func InitRoute(app *iris.Application, db *gorm.DB, conf *viper.Viper) {
	validator := validator.NewValidator()
	ur := repository.NewUserRepositoryImpl(db)
	us := service.NewUserServiceImpl(ur)
	uc := controller.NewUserController(us, validator)

	app.Get("/ping", Ping)

	app.Post("/signup", uc.SignUp)
	app.Post("/login", uc.LogIn)
}

func Ping(ctx iris.Context) {
	ctx.JSON(response.NewBaseResponse("pong", nil, nil))
}
