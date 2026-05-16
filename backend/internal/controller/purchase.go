package controller

import (
	"erp/internal/dto"
	"erp/internal/middleware"
	"erp/internal/response"
	"erp/internal/service"

	"github.com/gin-gonic/gin"
)

type PurchaseController struct {
	svc *service.PurchaseService
}

func NewPurchaseController(svc *service.PurchaseService) *PurchaseController {
	return &PurchaseController{svc: svc}
}

func (ctl *PurchaseController) CreateOrder(c *gin.Context) {
	var req dto.PurchaseOrderRequest
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

func (ctl *PurchaseController) UpdateOrder(c *gin.Context) {
	id, ok := ParseID(c)
	if !ok {
		return
	}
	var req dto.PurchaseOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.BadRequest("请求参数错误"))
		return
	}
	if err := ctl.svc.UpdateOrder(id, req); err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, gin.H{"id": id})
}

func (ctl *PurchaseController) ApproveOrder(c *gin.Context) { ctl.action(c, ctl.svc.ApproveOrder) }
func (ctl *PurchaseController) CloseOrder(c *gin.Context)   { ctl.action(c, ctl.svc.CloseOrder) }
func (ctl *PurchaseController) VoidOrder(c *gin.Context)    { ctl.action(c, ctl.svc.VoidOrder) }
func (ctl *PurchaseController) VoidInbound(c *gin.Context)  { ctl.action(c, ctl.svc.VoidInbound) }

func (ctl *PurchaseController) CreateInbound(c *gin.Context) {
	var req dto.PurchaseInboundRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.BadRequest("请求参数错误"))
		return
	}
	inbound, err := ctl.svc.CreateInbound(req)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Created(c, inbound)
}

func (ctl *PurchaseController) ConfirmInbound(c *gin.Context) {
	id, ok := ParseID(c)
	if !ok {
		return
	}
	if err := ctl.svc.ConfirmInbound(id, middleware.UserID(c)); err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, gin.H{"id": id})
}

func (ctl *PurchaseController) action(c *gin.Context, fn func(uint) error) {
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
