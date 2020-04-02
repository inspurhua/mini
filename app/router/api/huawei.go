package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/service"
	"huage.tech/mini/app/util"
)

func DeviceList(c *gin.Context) {
	data, err := service.GetDevices(0, 50)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(data, 1))
}


func DeviceCommand(c *gin.Context) {
	deviceId := c.Param("device_id")
	if deviceId == "" {
		util.AbortNewResultErrorOfClient(c, errors.New("需要传入设备id"))
		return
	}
	var form service.HwCommand
	err := c.ShouldBind(&form)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	result, err := service.SendDeviceCommand(deviceId, form)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite(result, 1))

}
