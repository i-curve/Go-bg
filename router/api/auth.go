package api

import (
	"net/http"
	"template/pkg/app"
	"template/pkg/e"
	"template/pkg/util"

	"template/service/auth_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	var a auth
	a.Username = c.PostForm("username")
	a.Password = c.PostForm("password")
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MakeErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, valid.Errors)
		return
	}

	authService := auth_service.Auth{Username: a.Username, Password: a.Password}
	isExist := authService.IsExist()
	if isExist == false {
		appG.Response(http.StatusOK, e.USER_NOT_EXIST, nil)
		return
	}
	isCheck, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	if !isCheck {
		appG.Response(http.StatusOK, e.USER_PASSWORD_ERROR, nil)
		return
	}

	token, err := util.GenerateToken(a.Username, a.Password)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_GENERNATE_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	var a auth
	a.Username = c.PostForm("username")
	a.Password = c.PostForm("password")
	ok, _ := valid.Valid(&a)
	if !ok {
		app.MakeErrors(valid.Errors)
		appG.Response(http.StatusAccepted, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: a.Username, Password: a.Password}
	isExist := authService.IsExist()
	if isExist == true {
		appG.Response(http.StatusOK, e.USER_ALREADY_EXIST, nil)
		return
	}
	status, err := authService.Register()
	if status != true || err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"status": "ok",
	})
}
