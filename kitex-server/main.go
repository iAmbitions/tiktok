package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	comment "tiktok/kitex-server/kitex_gen/douyin/comment/internalcommentservice"
	favorite "tiktok/kitex-server/kitex_gen/douyin/favorite/internalfavoriteservice"
)

func main() {
	svr := favorite.NewServer(NewFavoriteactionService())

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8989")
	if err != nil {
		panic(err)
	}
	svr1 := comment.NewServer(NewCommentactionService(),
		server.WithServiceAddr(addr))
	err1 := svr1.Run()

	if err1 != nil {
		log.Println(err1.Error())
	}
}
