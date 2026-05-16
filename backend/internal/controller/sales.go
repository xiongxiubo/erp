package controller

import (
	"erp/internal/dto"
	"erp/internal/middleware"
	"erp/internal/response"
	"erp/internal/service"

	"github.com/gin-gonic/gin"
)

type SalesController struct {
	svc *service.SalesService
}

func NewSalesController(svc *service.SalesService) *SalesController {
	return &SalesController{svc: svc}
}

func (ctl *SalesController) CreateOrder(c *gin.Context) {
	var req dto.SalesOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.BadRequest("请求参数错误"))
		return
	}
	order, err := ctl.svc.CreateOrder(req)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Created(c, order)
}

func (ctl *SalesController) ApproveOrder(c *gin.Context) { ctl.action(c, ctl.svc.ApproveOrder) }
func (ctl *SalesController) CloseOrder(c *gin.Context)   { ctl.action(c, ctl.svc.CloseOrder) }
func (ctl *SalesController) VoidOrder(c *gin.Context)    { ctl.action(c, ctl.svc.VoidOrder) }
func (ctl *SalesController) VoidOutbound(c *gin.Context) { ctl.action(c, ctl.svc.VoidOutbound) }

func (ctl *SalesController) CreateOutbound(c *gin.Context) {
	var req dto.SalesOutboundRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.BadRequest("请求参数错误"))
		return
	}
	outbound, err := ctl.svc.CreateOutbound(req)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Created(c, outbound)
}

func (ctl *SalesController) ConfirmOutbound(c *gin.Context) {
	id, ok := ParseID(c)
	if !ok {
		return
	}
	if err := ctl.svc.ConfirmOutbound(id, middleware.UserID(c)); err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, gin.H{"id": id})
}

func (ctl *SalesController) action(c *gin.Context, fn func(uint) error) {
	id, ok := ParseID(c)
	if !ok {
		return
	}
	if err := fn(id); err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, gin.H{"id": id})
}
