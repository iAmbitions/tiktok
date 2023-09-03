// Code generated by Kitex v0.7.0. DO NOT EDIT.

package feedservice

import (
	server "github.com/cloudwego/kitex/server"
	feed "tiktok/hertz-server/kitex_gen/douyin/feed"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler feed.FeedService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
