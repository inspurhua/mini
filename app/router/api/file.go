package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"os"
	"strconv"
	"syscall"
	"time"
)

func FileCreate(c *gin.Context) {
	var err error
	file, err := c.FormFile("file")
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	urlPrefix := "/uploads/" + time.Now().Format("20060102") + "/"
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
