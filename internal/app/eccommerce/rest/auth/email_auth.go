package auth

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	auth_service "eccomerce_apis/internal/app/eccommerce/services/auth"
	"eccomerce_apis/pkg/middlewares"
	"os"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

func (api *AuthApi) authorize_email(ctx iris.Context) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var auth models.UserAuth
	var res = make(map[string]interface{})
	var body Body
	err := ctx.ReadJSON(&body)
	if err != nil {
		ctx.StatusCode(iris.StatusOK)
		res["error"] = "unable to read body"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	if len(body.Email) == 0 || len(body.Password) == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "make sure both the email and password exists"
		ctx.JSON(res)
		return
	}

	err = api.Config.DB.Model(&auth).Where("email = ?", body.Email).Select()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to fetch user information"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ok := middlewares.CheckPassword(body.Password, auth.Password)
	if !ok {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "password is incorrect"
		ctx.JSON(res)
		return
	}

	claims := middlewares.JWTClaims{
		Role: "user",
		Id:   auth.UserID.String(),
	}
	token, err := claims.GenerateToken(os.Getenv("SIGNING_KEY"))
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to generate token"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	res["token"] = token
	ctx.JSON(res)
}

func (api *AuthApi) setup_email_password(ctx iris.Context) {
	type Body struct {
		UserID   uuid.UUID `json:"user_id"`
		Password string    `json:"password"`
	}

	var body Body
	var res = make(map[string]interface{})

	err := ctx.ReadJSON(&body)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to read body"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	var auth models.UserAuth

	if body.UserID == uuid.Nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "user id is not found"
		ctx.JSON(res)
		return
	}

	password, err := middlewares.HashPassword(body.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to generate password"
		ctx.JSON(res)
		return
	}

	auth.Password = password
	_, err = api.Config.DB.Model(&auth).Column("password").Where("user_id = ?", body.UserID).Update()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to fetch auth info"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	res["message"] = "password updated successfully"
	ctx.JSON(res)
}

func (api *AuthApi) register_with_email(ctx iris.Context) {
	var data auth_service.EmailPostData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to register read body",
			"error_info": err.Error(),
		})
		return
	}

	id, err := auth_service.RegisterWithEmail(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": "unable to register email",
		})
		return

	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"id": id,
	})
}

func (api *AuthApi) validate_with_email(ctx iris.Context) {
	var data auth_service.EmailPostData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to register read body",
			"error_info": err.Error(),
		})
		return
	}

	token, err := auth_service.ValidateEmailRegistration(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to register email",
			"error_info": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"token": token,
	})
}
