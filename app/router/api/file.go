package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unicode"
)

func FileCreate(c *gin.Context) {
	var err error
	file, err := c.FormFile("file")
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	//存储类型,只能是英文字符,common or tmp
	StoreType := strings.Trim(c.DefaultPostForm("store", ""), " ")
	if StoreType == "" || !(strings.Contains("<tmp><common>", "<"+StoreType+">")) {
		util.AbortNewResultErrorOfClient(c, errors.New("请提供store参数,tmp为临时存储,common为永久存储"))
		return
	}
	StoreType = StoreType + "/"
	//业务文件类型,about a业务还是b业务
	AboutType := strings.Trim(c.DefaultPostForm("about", "default"), " ")
	if AboutType != "" {
		for _, r := range AboutType {
			if !unicode.IsLetter(r) {
				util.AbortNewResultErrorOfClient(c, err)
				return
			}
		}
		AboutType = AboutType + "/"
	}

	urlPrefix := "/uploads/" + StoreType + AboutType + time.Now().Format("20060102") + "/"
	path := "./public/dist" + urlPrefix
	syscall.Umask(0)
	os.MkdirAll(path, 0777)
	err = c.SaveUploadedFile(file, path+file.Filename)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	form := bean.File{
		ID:       0,
		Name:     file.Filename,
		SaveName: path + file.Filename,
		SavePath: path,
		Url:      urlPrefix + file.Filename,
		CreateAt: time.Now(),
	}

	r, err := dao.FileCreate(form)

	c.JSON(200, util.NewResultOKofWrite(r, 1))
	return
}

func FileRead(c *gin.Context) {
	roleId := c.Param("id")
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	r, err := dao.FileRead(id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}
