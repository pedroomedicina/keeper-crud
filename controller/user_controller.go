package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"keeper-crud/data/request"
	"keeper-crud/data/response"
	"keeper-crud/helper"
	"keeper-crud/service"
	"net/http"
)

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{usersService: service}
}

// Signup SignUp Users		godoc
//
//	@Summary		SignUp users
//	@Description	Save users data in Db.
//	@Param			users	body	request.UserSignUpRequest	true	"Signup users"
//	@Produce		application/json
//	@Users			users
//	@Success		200	{object}	response.Response{}
//	@Router			/signup [post]
func (controller *UsersController) Signup(ctx *gin.Context) {
	log.Info().Msg("signup user")
	userSignUpRequest := request.UserSignUpRequest{}
	err := ctx.ShouldBindJSON(&userSignUpRequest)
	helper.ErrorPanic(err)

	err = controller.usersService.SignUp(userSignUpRequest)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Signin Sign In Users      godoc
//
//	@Summary		Sign In users
//	@Description	Authenticate users and create a session.
//	@Param			loginDetails	body	request.UserSignInRequest	true	"Signin user details"
//	@Produce		application/json
//	@Tags			users
//	@Success		200	{object}	response.Response{}
//	@Failure		400	{object}	response.ErrorResponse{}
//	@Failure		401	{object}	response.ErrorResponse{}
//	@Failure		500	{object}	response.ErrorResponse{}
//	@Router			/signin [post]
func (controller *UsersController) Signin(ctx *gin.Context) {
	log.Info().Msg("signin user")
	session := sessions.Default(ctx)
	loginDetails := request.UserSignInRequest{}

	if err := ctx.ShouldBindJSON(&loginDetails); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request"})
		return
	}

	user, err := controller.usersService.AuthenticateUser(loginDetails.Email, loginDetails.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Error: "Authentication failed"})
		return
	}

	session.Set("user_id", user.ID)
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to save session"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User signed in successfully"})
}
