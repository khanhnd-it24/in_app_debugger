package controllers

import (
	"backend/src/core/services"
	"backend/src/present/http/requests"
	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	*baseController
	deviceService *services.DeviceService
}

func NewDeviceController(
	baseController *baseController,
	deviceService *services.DeviceService,
) *DeviceController {
	return &DeviceController{
		baseController: baseController,
		deviceService:  deviceService,
	}
}

func (i *DeviceController) UpdateDevice(c *gin.Context) {
	req := new(requests.Device)
	if err := i.BindAndValidateRequest(c, req); err != nil {
		i.ErrorData(c, err)
		return
	}

	device, ierr := i.deviceService.UpsertDeviceUC(c, req)
	if ierr != nil {
		i.ErrorData(c, ierr)
		return
	}

	i.Success(c, device)
}

func (i *DeviceController) GetDevice(c *gin.Context) {
	deviceId := c.Param("device_id")
	device, ierr := i.deviceService.GetDeviceByDeviceId(c, deviceId)
	if ierr != nil {
		i.ErrorData(c, ierr)
		return
	}
	i.Success(c, device)
}
