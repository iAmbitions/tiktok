// Code generated by Kitex v0.7.0. DO NOT EDIT.

package internalcommentservice

import (
			"context"
				client "github.com/cloudwego/kitex/client"
				kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
				comment "tiktok/hertz-server/kitex_gen/douyin/comment"
)

func serviceInfo() *kitex.ServiceInfo {
	return internalCommentServiceServiceInfo
 }

var internalCommentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "InternalCommentService"
	handlerType := (*comment.InternalCommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"InternalActionComment":
			kitex.NewMethodInfo(internalActionCommentHandler, newInternalCommentServiceInternalActionCommentArgs, newInternalCommentServiceInternalActionCommentResult, false),
		"InternalListComment":
			kitex.NewMethodInfo(internalListCommentHandler, newInternalCommentServiceInternalListCommentArgs, newInternalCommentServiceInternalListCommentResult, false),
		"InternalCountComment":
			kitex.NewMethodInfo(internalCountCommentHandler, newInternalCommentServiceInternalCountCommentArgs, newInternalCommentServiceInternalCountCommentResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":	 "comment",
		"ServiceFilePath": "..\idl\comment.thrift",
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



func internalActionCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*comment.InternalCommentServiceInternalActionCommentArgs)
	realResult := result.(*comment.InternalCommentServiceInternalActionCommentResult)
	success, err := handler.(comment.InternalCommentService).InternalActionComment(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newInternalCommentServiceInternalActionCommentArgs() interface{} {
	return comment.NewInternalCommentServiceInternalActionCommentArgs()
}

func newInternalCommentServiceInternalActionCommentResult() interface{} {
	return comment.NewInternalCommentServiceInternalActionCommentResult()
}


func internalListCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*comment.InternalCommentServiceInternalListCommentArgs)
	realResult := result.(*comment.InternalCommentServiceInternalListCommentResult)
	success, err := handler.(comment.InternalCommentService).InternalListComment(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newInternalCommentServiceInternalListCommentArgs() interface{} {
	return comment.NewInternalCommentServiceInternalListCommentArgs()
}

func newInternalCommentServiceInternalListCommentResult() interface{} {
	return comment.NewInternalCommentServiceInternalListCommentResult()
}


func internalCountCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*comment.InternalCommentServiceInternalCountCommentArgs)
	realResult := result.(*comment.InternalCommentServiceInternalCountCommentResult)
	success, err := handler.(comment.InternalCommentService).InternalCountComment(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newInternalCommentServiceInternalCountCommentArgs() interface{} {
	return comment.NewInternalCommentServiceInternalCountCommentArgs()
}

func newInternalCommentServiceInternalCountCommentResult() interface{} {
	return comment.NewInternalCommentServiceInternalCountCommentResult()
}


type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}


func (p *kClient) InternalActionComment(ctx context.Context , request *comment.InternalActionCommentRequest) (r *comment.InternalActionCommentResponse, err error) {
	var _args comment.InternalCommentServiceInternalActionCommentArgs
	_args.Request = request
	var _result comment.InternalCommentServiceInternalActionCommentResult
	if err = p.c.Call(ctx, "InternalActionComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) InternalListComment(ctx context.Context , request *comment.InternalListCommentRequest) (r *comment.InternalListCommentResponse, err error) {
	var _args comment.InternalCommentServiceInternalListCommentArgs
	_args.Request = request
	var _result comment.InternalCommentServiceInternalListCommentResult
	if err = p.c.Call(ctx, "InternalListComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) InternalCountComment(ctx context.Context , request *comment.InternalCountCommentRequest) (r *comment.InternalCountCommentResponse, err error) {
	var _args comment.InternalCommentServiceInternalCountCommentArgs
	_args.Request = request
	var _result comment.InternalCommentServiceInternalCountCommentResult
	if err = p.c.Call(ctx, "InternalCountComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

