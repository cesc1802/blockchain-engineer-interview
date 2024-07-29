package main

import (
	"backend-golang/component/appctx"
	tokenbizs "backend-golang/modules/token/biz"
	tokengin "backend-golang/modules/token/transport/gin"

	"github.com/gin-gonic/gin"
)

func Route(g *gin.RouterGroup, appCtx appctx.AppContext, tokenBiz tokenbizs.TokenBiz) {
	tokens := g.Group("/tokens")
	{
		tokens.POST("/savesession", tokengin.CreateToken(appCtx, tokenBiz))
		tokens.POST("/confirm", tokengin.Confirm(appCtx, tokenBiz))
		tokens.GET("/uploadsessions", tokengin.GetListUploadSessions(appCtx, tokenBiz))
	}
}
