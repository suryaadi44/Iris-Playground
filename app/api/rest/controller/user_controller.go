package controller

import (
	"github.com/suryaadi44/iris-playground/app/dto"
	"github.com/suryaadi44/iris-playground/app/usecase"
	"github.com/suryaadi44/iris-playground/utils/response"
	"github.com/suryaadi44/iris-playground/utils/validator"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	us        usecase.UserService
	validator validator.Validator
}

func NewUserController(us usecase.UserService, validator validator.Validator) *UserController {
	return &UserController{
		us:        us,
		validator: validator,
	}
}

func (c *UserController) SignUp(ctx iris.Context) {
	req := new(dto.UserSignUpRequest)
	if err := ctx.ReadJSON(req); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest,
			response.NewErrorResponse(
				response.ResponseInvalidRequest,
				*response.NewErrorValue("request body", err.Error()),
			),
		)
		return
	}

	if errs := c.validator.ValidateJSON(req); errs != nil {
		ctx.StopWithJSON(iris.StatusBadRequest,
			response.NewBaseResponse(response.ResponseValidationError, nil, *errs))
		return
	}

	err := c.us.SignUp(ctx.Request().Context(), req)
	if err != nil {
		switch err {
		case response.ErrDuplicateEmail:
			ctx.StopWithJSON(iris.StatusConflict,
				response.NewErrorResponse(
					response.ResponseBusinessLogicError,
					*response.NewErrorValue("email", err.Error()),
				),
			)
		default:
			ctx.StopWithJSON(iris.StatusInternalServerError,
				response.NewErrorResponse(
					response.ResponseRuntimeError,
					*response.NewErrorValue("server error", err.Error()),
				),
			)
		}
		return
	}

	ctx.JSON(response.NewBaseResponse(response.ResponseSuccess, nil, nil))
}

func (c *UserController) LogIn(ctx iris.Context) {
	req := new(dto.UserLoginRequest)
	if err := ctx.ReadJSON(req); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest,
			response.NewErrorResponse(
				response.ResponseInvalidRequest,
				*response.NewErrorValue("request body", err.Error()),
			),
		)
		return
	}

	if errs := c.validator.ValidateJSON(req); errs != nil {
		ctx.StopWithJSON(iris.StatusBadRequest,
			response.NewBaseResponse(response.ResponseValidationError, nil, *errs))
		return
	}

	token, err := c.us.LogIn(ctx.Request().Context(), req)
	if err != nil {
		switch err {
		case response.ErrInvalidEmailOrPassword:
			ctx.StopWithJSON(iris.StatusUnauthorized,
				response.NewErrorResponse(
					response.ResponseBusinessLogicError,
					*response.NewErrorValue("email or password", err.Error()),
				),
			)
		default:
			ctx.StopWithJSON(iris.StatusInternalServerError,
				response.NewErrorResponse(
					response.ResponseRuntimeError,
					*response.NewErrorValue("server error", err.Error()),
				),
			)
		}
		return
	}

	ctx.JSON(response.NewBaseResponse(
		response.ResponseSuccess,
		token,
		nil,
	))
}
