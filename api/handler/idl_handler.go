// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wuhuZhao/overpass_frontend_api/api/dao"
	"github.com/wuhuZhao/overpass_frontend_api/internal"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type IdlHandler struct {
	dao *dao.IdlDao
}

func NewIdlHandler(db *gorm.DB) *IdlHandler {
	return &IdlHandler{dao: dao.NewIdlDao(db)}
}

// Create .
func (handler *IdlHandler) Create(ctx context.Context, c *app.RequestContext) {
	var idl *dao.Idl
	err := c.BindAndValidate(idl)
	if err != nil {
		c.JSON(http.StatusBadRequest, internal.Response{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	idl.CreateTime = time.Now()
	idl.Use = true
	idl.Tag = ""
	var preIdl *dao.Idl
	err = handler.dao.FindByName(preIdl, idl.Name)
	if err != nil {
		idl.Version = 0
	} else {
		idl.Version = preIdl.Version + 1
	}
	err = handler.dao.InsertOne(idl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, internal.Response{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, internal.Response{
		Code: 0,
		Msg:  "insert success",
		Data: idl,
	})
}

// Update .
func (handler *IdlHandler) Update(ctx context.Context, c *app.RequestContext) {

}

// Delete .
func (handler *IdlHandler) Delete(ctx context.Context, c *app.RequestContext) {

}

// Find .
func (handler *IdlHandler) Find(ctx context.Context, c *app.RequestContext) {

}

// FindAll .
func (handler *IdlHandler) FindAll(ctx context.Context, c *app.RequestContext) {

}
