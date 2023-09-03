// Code generated by Kitex v0.7.0. DO NOT EDIT.

package feedservice

import (
			"context"
				client "github.com/cloudwego/kitex/client"
				kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
				feed "tiktok/kitex-server/kitex_gen/douyin/feed"
)

func serviceInfo() *kitex.ServiceInfo {
	return feedServiceServiceInfo
 }

var feedServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FeedService"
	handlerType := (*feed.FeedService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ListVideos":
			kitex.NewMethodInfo(listVideosHandler, newFeedServiceListVideosArgs, newFeedServiceListVideosResult, false),
		"QueryVideos":
			kitex.NewMethodInfo(queryVideosHandler, newFeedServiceQueryVideosArgs, newFeedServiceQueryVideosResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":	 "feed",
		"ServiceFilePath": "..\IDL\feed.thrift",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName: 	 serviceName,
		HandlerType: 	 handlerType,
		Methods:     	 methods,
		PayloadCodec:  	 kitex.Thrift,
		KiteXGenVersion: "v0.7.0",
		Extra:           extra,
	}
	return svcInfo
}



func listVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*feed.FeedServiceListVideosArgs)
	realResult := result.(*feed.FeedServiceListVideosResult)
	success, err := handler.(feed.FeedService).ListVideos(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newFeedServiceListVideosArgs() interface{} {
	return feed.NewFeedServiceListVideosArgs()
}

func newFeedServiceListVideosResult() interface{} {
	return feed.NewFeedServiceListVideosResult()
}


func queryVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*feed.FeedServiceQueryVideosArgs)
	realResult := result.(*feed.FeedServiceQueryVideosResult)
	success, err := handler.(feed.FeedService).QueryVideos(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newFeedServiceQueryVideosArgs() interface{} {
	return feed.NewFeedServiceQueryVideosArgs()
}

func newFeedServiceQueryVideosResult() interface{} {
	return feed.NewFeedServiceQueryVideosResult()
}


type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}


func (p *kClient) ListVideos(ctx context.Context , request *feed.ListFeedRequest) (r *feed.ListFeedResponse, err error) {
	var _args feed.FeedServiceListVideosArgs
	_args.Request = request
	var _result feed.FeedServiceListVideosResult
	if err = p.c.Call(ctx, "ListVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryVideos(ctx context.Context , request *feed.QueryVideosRequest) (r *feed.QueryVideosResponse, err error) {
	var _args feed.FeedServiceQueryVideosArgs
	_args.Request = request
	var _result feed.FeedServiceQueryVideosResult
	if err = p.c.Call(ctx, "QueryVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

