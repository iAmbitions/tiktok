// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	douyin_api_favorite "tiktok/hertz-server/biz/router/douyin/api/favorite"
	douyin_comment "tiktok/hertz-server/biz/router/douyin/comment"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	douyin_comment.Register(r)

	douyin_api_favorite.Register(r)

}
