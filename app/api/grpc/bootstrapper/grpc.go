package bootstrapper

import (
	"github.com/spf13/viper"
	"github.com/suryaadi44/iris-playground/app/api/grpc/server"
	repository "github.com/suryaadi44/iris-playground/app/repository/impl"
	usecase "github.com/suryaadi44/iris-playground/app/usecase/impl"
	"github.com/suryaadi44/iris-playground/utils/validator"
	"gorm.io/gorm"
)

func InitGRPC(db *gorm.DB, conf *viper.Viper) *server.AuthServer {
	validator := validator.NewValidator()
	ur := repository.NewUserRepositoryImpl(db)
	us := usecase.NewUserServiceImpl(ur)
	s := server.NewAuthServer(us, validator)

	return s
}
