package controllers

import (
	"backend/src/core/services"
	"backend/src/present/http/requests"
	"github.com/gin-gonic/gin"
)

type ConsoleController struct {
	*baseController
	consoleService *services.ConsoleService
}

func NewConsoleController(
	baseController *baseController,
	consoleService *services.ConsoleService,
) *ConsoleController {
	return &ConsoleController{
		baseController: baseController,
		consoleService: consoleService,
	}
}

func (con *ConsoleController) CreateConsole(c *gin.Context) {
	req := new(requests.Console)
	if err := con.BindAndValidateRequest(c, req); err != nil {
		con.ErrorData(c, err)
		return
	}

	console, ierr := con.consoleService.CreateConsoleUC(c, req)
	if ierr != nil {
		con.ErrorData(c, ierr)
		return
	}

	con.Success(c, console)
}

func (con *ConsoleController) GetConsoles(c *gin.Context) {
	deviceId := c.Param("device_id")
	consoles, ierr := con.consoleService.GetConsolesByDeviceId(c, deviceId)
	if ierr != nil {
		con.ErrorData(c, ierr)
		return
	}

	con.Success(c, consoles)
}
