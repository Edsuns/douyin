package main

import (
	"douyin/app/api"
	"douyin/app/config"
	"douyin/pkg/security"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	initRouter(r)
	return r
}

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static(config.Val.Static.Route, config.Val.Static.Filepath)

	apiRouter := r.Group("/douyin")

	/* public apis */
	apiRouter.POST("/user/register/", api.Register)
	apiRouter.POST("/user/login/", api.Login)

	/* security apis */
	apiRouter.Use(security.Middleware)
	{
		// basic apis
		apiRouter.GET("/feed/", api.Feed)
		apiRouter.GET("/user/", api.UserInfo)
		apiRouter.POST("/publish/action/", api.Publish)
		apiRouter.GET("/publish/list/", api.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", api.FavoriteAction)
		apiRouter.GET("/favorite/list/", api.FavoriteList)
		apiRouter.POST("/comment/action/", api.CommentAction)
		apiRouter.GET("/comment/list/", api.CommentList)

		// extra apis - II
		apiRouter.POST("/relation/action/", api.RelationAction)
		apiRouter.GET("/relation/follow/list/", api.FollowList)
		apiRouter.GET("/relation/follower/list/", api.FollowerList)
	}
}
