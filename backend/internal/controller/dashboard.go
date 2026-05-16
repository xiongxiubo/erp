package controller

import (
	"erp/internal/response"
	"erp/internal/service"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	svc *service.DashboardService
}

func NewDashboardController(svc *service.DashboardService) *DashboardController {
	return &DashboardController{svc: svc}
}

func (ctl *DashboardController) Summary(c *gin.Context) {
	data, err := ctl.svc.Summary()
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, data)
}

func (ctl *DashboardController) RecentDocuments(c *gin.Context) {
	data, err := ctl.svc.RecentDocuments()
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, data)
}
