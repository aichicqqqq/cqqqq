package routes

import (
	"todolist/api"
	"todolist/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 配置路由
func Newrouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something secret"))
	r.Use(sessions.Sessions("mysession", store))
	//r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())
		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			// 任务操作
			authed.POST("task_create", api.CreateTaskHandler())
			authed.GET("task_list", api.ListTaskHandler())
			authed.GET("task_show", api.ShowTaskHandler())
			authed.POST("task_update", api.UpdateTaskHandler())
			authed.POST("task_search", api.SearchTaskHandler())
			authed.POST("task_delete", api.DeleteTaskHandler())
		}
	}
	return r

}
