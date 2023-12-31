// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "tiktok/hertz-server/biz/handler"
	favorite "tiktok/hertz-server/biz/handler/douyin/api/favorite"
	comment "tiktok/hertz-server/biz/handler/douyin/comment"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.POST("/douyin/favorite/action/", favorite.FavoriteAction)
	r.GET("/douyin/favorite/list/", favorite.FavoriteList)
	r.POST("/douyin/comment/action/", comment.ActionComment)
	r.GET("/douyin/comment/list/", comment.ListComment)
	// your code ...
}
