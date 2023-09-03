// Code generated by hertz generator.

package comment

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	client2 "github.com/cloudwego/kitex/client"
	comment "tiktok/hertz-server/biz/model/douyin/comment"
	comment1 "tiktok/hertz-server/kitex_gen/douyin/comment"
	"tiktok/hertz-server/kitex_gen/douyin/comment/internalcommentservice"
	"tiktok/mw/jwt"
)

func convertInternalCommentToComment(internalComment *comment1.InternalComment) *comment.Comment {
	return &comment.Comment{
		ID:         internalComment.Id,
		UserID:     internalComment.UserId,
		Content:    internalComment.Content,
		CreateDate: internalComment.CreateDate,
	}
}

func ActionComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.ActionCommentRequest
	// Parameter binding and validation capabilities provided by hertz
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	client, err := internalcommentservice.NewClient("comment", client2.WithHostPorts("127.0.0.1:8989"))
	if err != nil {
		panic(err)
	}
	loginInfo, err := jwt.Parse(req.Token)

	reqRpc := &comment1.InternalActionCommentRequest{
		UserId:      loginInfo.UerID,
		VideoId:     req.VideoID,
		ActionType:  comment1.InternalActionCommentType(req.ActionType),
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	}

	respRpc, err := client.InternalActionComment(ctx, reqRpc)
	if err != nil {
		panic(err)
	}
	if respRpc.StatusCode != 0 {
		resp := &comment.ActionCommentResponse{
			StatusCode: respRpc.StatusCode,
			StatusMsg:  respRpc.StatusMsg,
			Comment:    convertInternalCommentToComment(respRpc.Comment),
		}
		c.JSON(200, resp)
		return
	}
	resp := &comment.ActionCommentResponse{
		StatusCode: respRpc.StatusCode,
		StatusMsg:  respRpc.StatusMsg,
		Comment:    convertInternalCommentToComment(respRpc.Comment),
	}
	c.JSON(200, resp)
}
func convertInternalCommentsToComments(internalComments []*comment1.InternalComment) []*comment.Comment {
	comments := make([]*comment.Comment, len(internalComments))
	for i, internalComment := range internalComments {
		comments[i] = convertInternalCommentToComment(internalComment)
	}
	return comments
}

func ListComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.ListCommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	client, err := internalcommentservice.NewClient("comment", client2.WithHostPorts("127.0.0.1:8989"))
	if err != nil {
		panic(err)
	}
	loginInfo, err := jwt.Parse(req.Token)
	reqRpc := &comment1.InternalListCommentRequest{
		UserId:  loginInfo.UserID,
		VideoId: req.VideoID,
	}
	respRpc, err := client.InternalListComment(ctx, reqRpc)
	if err != nil {
		panic(err)
	}
	if respRpc.StatusCode != 0 {
		resp := &comment.ListCommentResponse{
			StatusCode:  respRpc.StatusCode,
			StatusMsg:   respRpc.StatusMsg,
			CommentList: convertInternalCommentsToComments(respRpc.CommentList),
		}
		c.JSON(200, resp)
		return
	}
	resp := &comment.ListCommentResponse{
		StatusCode:  respRpc.StatusCode,
		StatusMsg:   respRpc.StatusMsg,
		CommentList: convertInternalCommentsToComments(respRpc.CommentList),
	}
	c.JSON(200, resp)
}
