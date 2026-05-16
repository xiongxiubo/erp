package controller

import (
	"erp/internal/response"
	"erp/internal/service"
	"erp/internal/utils"

	"github.com/gin-gonic/gin"
)

type CRUDController[T any] struct {
	svc *service.CRUDService[T]
}

func NewCRUDController[T any](svc *service.CRUDService[T]) *CRUDController[T] {
	return &CRUDController[T]{svc: svc}
}

func (ctl *CRUDController[T]) List(c *gin.Context) {
	p := utils.GetPagination(c)
	items, total, err := ctl.svc.List(p)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, response.Page{Items: items, Total: total, Page: p.Page, PageSize: p.PageSize})
}

func (ctl *CRUDController[T]) Get(c *gin.Context) {
	id, ok := ParseID(c)
	if !ok {
		return
	}
	item, err := ctl.svc.Get(id)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, item)
}

func (ctl *CRUDController[T]) Create(c *gin.Context) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		response.Fail(c, response.BadRequest("请求参数错误"))
		return
	}
	if err := ctl.svc.Create(&item); err != nil {
		response.Fail(c, err)
		return
	}
	response.Created(c, item)
}

func (ctl *CRUDController[T]) Update(c *gin.Context) {
	id, ok := ParseID(c)
	if !ok {
		return
	}
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		response.Fail(c, response.BadRequest("请求参数错误"))
		return
	}
	if err := ctl.svc.Update(id, &item); err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, gin.H{"id": id})
}

func (ctl *CRUDController[T]) Delete(c *gin.Context) {
	id, ok := ParseID(c)
	if !ok {
		return
	}
	if err := ctl.svc.Delete(id); err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, gin.H{"id": id})
}
