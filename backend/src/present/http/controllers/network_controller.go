package controllers

import (
	"backend/src/core/services"
	"backend/src/present/http/requests"
	"github.com/gin-gonic/gin"
)

type NetworkController struct {
	*baseController
	networkService *services.NetworkService
}

func NewNetworkController(
	baseController *baseController,
	networkService *services.NetworkService,
) *NetworkController {
	return &NetworkController{
		baseController: baseController,
		networkService: networkService,
	}
}

func (n *NetworkController) CreateNetwork(c *gin.Context) {
	req := new(requests.Network)
	if err := n.BindAndValidateRequest(c, req); err != nil {
		n.ErrorData(c, err)
		return
	}

	network, ierr := n.networkService.CreateNetworkUC(c, req)
	if ierr != nil {
		n.ErrorData(c, ierr)
		return
	}

	n.Success(c, network)
}

func (n *NetworkController) GetNetworks(c *gin.Context) {
	deviceId := c.Param("device_id")
	networks, ierr := n.networkService.GetNetworksByDeviceId(c, deviceId)
	if ierr != nil {
		n.ErrorData(c, ierr)
		return
	}
	n.Success(c, networks)
}
