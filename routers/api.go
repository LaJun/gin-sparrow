package routers

import (
	"gin/controller"
	"github.com/gin-gonic/gin"
)

func LoadApiRouters(e *gin.Engine) {
	//用户管理
	userController := new (controller.UserController)
	userGroup := e.Group("/user")
	{
		userGroup.GET("/add", userController.Add)
	}

	//文章类目管理
	categoryController := new (controller.UserController)
	categoryGroup := e.Group("/user")
	{
		categoryGroup.GET("/add", categoryController.Add)
	}

	//文章管理
	ArticleController := new (controller.UserController)
	articleGroup := e.Group("/user")
	{
		articleGroup.GET("/add",ArticleController.Add)
	}
}
