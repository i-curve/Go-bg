package v1

import (
	"log"
	"net/http"
	"template/pkg/app"
	"template/pkg/e"

	"template/service/diary_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type diary struct {
	ID       int    `json:"id" binding:"-"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Text     string `json:"text"`
}

func CreateDiary(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	var d diary
	_ = c.BindJSON(&d)
	ok, _ := valid.Valid(&d)
	log.Println(d)
	if !ok {
		app.MakeErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	diaryServer := diary_service.Diary{Username: d.Username, Title: d.Title, Text: d.Text}
	status := diaryServer.CreateDiary()
	if !status {
		appG.Response(http.StatusOK, e.ERROR_ADD_ARTICLE_FAIL, "创建失败")
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
func ModifyDiary(c *gin.Context) {
	appG := app.Gin{C: c}
	var d diary
	err := c.BindJSON(&d)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	diaryServer := diary_service.Diary{ID: d.ID, Title: d.Title, Text: d.Text}
	ok := diaryServer.ModifyDiary()
	if !ok {
		appG.Response(http.StatusOK, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
func GetDiarys(c *gin.Context) {
	appG := app.Gin{C: c}

	var username string
	username = c.PostForm("username")
	if username == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	diary := diary_service.Diary{Username: username}
	diarys, err := diary.GetDiarys(0, 10)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, "获取失败")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, diarys)
}
func GetDiaryCount(c *gin.Context) {
	appG := app.Gin{C: c}

	var username string
	username = c.PostForm("username")
	if username == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	diaryServer := diary_service.Diary{Username: username}
	count, err := diaryServer.GetDiaryCount()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]int{
		"count": count,
	})
}
