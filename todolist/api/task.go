package api

import (
	"net/http"
	"todolist/pkg/util"
	"todolist/service"
	"todolist/types"

	"github.com/gin-gonic/gin"
)

const BasePageLimit = 15

// 创建任务
func CreateTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreateTaskReq
		if err := c.ShouldBind(&req); err == nil {
			//参数校验
			l := service.GetTaskSrv()
			resp, err := l.CreateTask(c.Request.Context(), &req)
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

// 获取任务列表
func ListTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListTasksReq
		if err := c.ShouldBind(&req); err == nil {
			//参数校验
			if req.Limit == 0 {
				req.Limit = BasePageLimit
			}
			l := service.GetTaskSrv()
			resp, err := l.ListTask(c.Request.Context(), &req)
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

// 展示任务
func ShowTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowTaskReq
		if err := c.ShouldBind(&req); err == nil {
			//参数校验
			l := service.GetTaskSrv()
			resp, err := l.ShowTask(c.Request.Context(), &req)
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

// 查询任务
func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.SearchTask(ctx.Request.Context(), &req)
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

// 修改任务
func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(types.UpdateTaskReq)
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.UpdateTask(ctx.Request.Context(), req)
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

// 删除任务
func DeleteTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.DeleteTask(ctx.Request.Context(), &req)
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
