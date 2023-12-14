package api

import (
	"net/http"
	"todolist/pkg/util"
	"todolist/service"
	"todolist/types"

	"github.com/gin-gonic/gin"
)
//用户注册
func UserRegisterHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		var req types.UserServiceReq
		if err := c.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.Register(c.Request.Context(), &req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			c.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}
//用户登录
func UserLoginHandler()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.Login(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}
