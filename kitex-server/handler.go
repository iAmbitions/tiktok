package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"tiktok/kitex-server/database/mysql"
	comment "tiktok/kitex-server/kitex_gen/douyin/comment"
	favorite "tiktok/kitex-server/kitex_gen/douyin/favorite"
	"tiktok/kitex-server/kitex_gen/douyin/feed"
	"time"
)

// InternalFavoriteServiceImpl implements the last service interface defined in the IDL.
type InternalFavoriteServiceImpl struct {
	db *gorm.DB
}
type InternalCommentServiceImpl struct {
	db *gorm.DB
}

func NewFavoriteactionService() *InternalFavoriteServiceImpl {
	return &InternalFavoriteServiceImpl{
		db: mysql.InitDB(),
	}
}
func NewCommentactionService() *InternalCommentServiceImpl {
	return &InternalCommentServiceImpl{
		db: mysql.InitDB(),
	}
}

// InternalFavoriteAction implements the InternalFavoriteServiceImpl interface.
// InternalFavoriteAction implements the InternalFavoriteServiceImpl interface.
func (s *InternalFavoriteServiceImpl) InternalFavoriteAction(ctx context.Context, request *favorite.InternalFavoriteRequest) (resp *favorite.InternalFavoriteResponse, err error) {
	// Check if the request is nil
	if request == nil {
		// Handle the error case: you might want to return an error or an appropriate response structure indicating the invalid request
		err = fmt.Errorf("request cannot be nil")
		return nil, err
	}

	// Get the video with the requested VideoID

	type Video struct {
		gorm.Model
		ID       uint `gorm:"primaryKey"`
		AuthorID uint `gorm:"foreignKey:AuthorID;references:User.User_id"` // 外键关联用户
		PlayURL  string
		Title    string
		CoverURL string

		FavoritedCount int64
		CommentCount   int64
	}
	var video Video
	result := s.db.First(&video, request.VideoId)
	if result.Error != nil {
		err = fmt.Errorf("video not found")
		return nil, err
	}
	// Create a string variable
	var msg string
	// Update FavoritedCount
	if request.ActionType == 1 {
		s.db.Model(&video).Update("FavoritedCount", gorm.Expr("FavoritedCount + ?", 1))
		msg = "点赞成功"
	} else if request.ActionType == 0 {
		s.db.Model(&video).Update("FavoritedCount", gorm.Expr("FavoritedCount - ?", 1))
		msg = "取消点赞"
	}

	// If request is not nil, create a response with StatusCode=1 and StatusMsg="1"
	resp = &favorite.InternalFavoriteResponse{
		StatusCode: 0,
		StatusMsg:  &msg, // Use the address of the string variable
	}
	return resp, nil
}

// InternalFavoriteList implements the InternalFavoriteServiceImpl interface.
// InternalFavoriteList implements the InternalFavoriteServiceImpl interface.
func (s *InternalFavoriteServiceImpl) InternalFavoriteList(ctx context.Context, request *favorite.InternalFavoriteListRequest) (resp *favorite.InternalFavoriteListResponse, err error) {
	// Check if the request is nil
	if request == nil {
		// Handle the error case: you might want to return an error or an appropriate response structure indicating the invalid request
		err = fmt.Errorf("request cannot be nil")
		return nil, err
	}
	type Favorite struct {
		ID        int64          `json:"id"`
		UserID    int64          `json:"user_id"`
		VideoID   int64          `json:"video_id"`
		CreatedAt time.Time      `json:"create_at"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
	}

	// 根据 UserID 查询所有收藏
	var favorites []Favorite
	if err := s.db.Where("UserID = ?", request.UserId).Find(&favorites).Error; err != nil {
		return nil, err
	}

	// 根据 favoriteList 中的 VideoID 查询所有视频
	var videoList []*feed.Video
	for _, fav := range favorites {
		var video feed.Video
		if err := s.db.Where("ID = ?", fav.VideoID).First(&video).Error; err != nil {
			return nil, err
		}
		videoList = append(videoList, &video)
	}

	msg := "1" // Create a string variable

	// Create the response object with StatusCode and StatusMsg
	resp = &favorite.InternalFavoriteListResponse{
		StatusCode: 0, // Assuming 0 means success
		StatusMsg:  &msg,
		VideoList:  videoList,
	}

	return resp, nil
}

// InternalIsFavorite implements the InternalFavoriteServiceImpl interface.
func (s *InternalFavoriteServiceImpl) InternalIsFavorite(ctx context.Context, request *favorite.InternalIsFavoriteRequest) (resp *favorite.InternalIsFavoriteResponse, err error) {
	// TODO: Your code here...
	if request == nil {
		// Handle the error case: you might want to return an error or an appropriate response structure indicating the invalid request
		err = fmt.Errorf("request cannot be nil")
		return nil, err
	}

	return
}

// InternalCountFavorite implements the InternalFavoriteServiceImpl interface.
func (s *InternalFavoriteServiceImpl) InternalCountFavorite(ctx context.Context, request *favorite.InternalCountFavoriteRequest) (resp *favorite.InternalCountFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// InternalCountUserFavorite implements the InternalFavoriteServiceImpl interface.
func (s *InternalFavoriteServiceImpl) InternalCountUserFavorite(ctx context.Context, request *favorite.InternalCountUserFavoriteRequest) (resp *favorite.InternalCountUserFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// InternalCountUserTotalFavorited implements the InternalFavoriteServiceImpl interface.
func (s *InternalFavoriteServiceImpl) InternalCountUserTotalFavorited(ctx context.Context, request *favorite.InternalCountUserTotalFavoritedRequest) (resp *favorite.InternalCountUserTotalFavoritedResponse, err error) {
	// TODO: Your code here...
	return
}

// InternalActionComment implements the InternalCommentServiceImpl interface.
func (s *InternalCommentServiceImpl) InternalActionComment(ctx context.Context, request *comment.InternalActionCommentRequest) (resp *comment.InternalActionCommentResponse, err error) {
	// TODO: Your code here...
	if request == nil {
		// Handle the error case: you might want to return an error or an appropriate response structure indicating the invalid request
		err = fmt.Errorf("request cannot be nil")
		return nil, err
	}
	resp = &comment.InternalActionCommentResponse{}
	type Comment struct {
		ID        int32     `gorm:"primaryKey"`
		VideoID   int32     `gorm:"column:video_id"`
		UserID    int32     `gorm:"column:user_id"`
		Content   string    `gorm:"column:content"`
		CreatedAt time.Time `gorm:"column:created_at"`
	}

	switch request.ActionType {
	case 1:
		comment1 := &Comment{
			VideoID:   request.VideoId,
			UserID:    request.UserId,
			Content:   *request.CommentText,
			CreatedAt: time.Now(),
		}
		if err := s.db.Create(comment1).Error; err != nil {
			resp.StatusCode = 500
			resp.StatusMsg = "Failed to add comment to database"
			return resp, err
		}

		resp.StatusCode = 200
		resp.StatusMsg = "Comment added successfully"
		resp.Comment = &comment.InternalComment{
			Id:         comment1.ID,
			UserId:     comment1.UserID,
			Content:    comment1.Content,
			CreateDate: comment1.CreatedAt.Format("2006-01-02 15:04:05"),
		}

	case 0:
		if err := s.db.Delete(&Comment{}, request.CommentId).Error; err != nil {
			resp.StatusCode = 500
			resp.StatusMsg = "Failed to delete comment from database"
			resp.Comment = nil
			return resp, err
		}
		resp.StatusCode = 200
		resp.StatusMsg = "Comment deleted successfully"
		resp.Comment = nil

	default:
		resp.StatusCode = 400
		resp.StatusMsg = "Invalid action_type"
		resp.Comment = nil
		return resp, fmt.Errorf("invalid action_type")
	}

	return resp, nil
}

func (s *InternalCommentServiceImpl) InternalListComment(ctx context.Context, request *comment.InternalListCommentRequest) (resp *comment.InternalListCommentResponse, err error) {
	if request == nil {
		err = fmt.Errorf("request cannot be nil")
		return nil, err
	}
	type Comment struct {
		ID         int64
		UserID     int64
		VideoID    int64
		Content    string
		CreateDate time.Time
	}

	var comments []Comment
	err = s.db.Where("video_id = ?", request.VideoId).Order("create_date desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	resp = &comment.InternalListCommentResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	res := []*comment.InternalComment{}
	for _, comment1 := range comments {
		createDate := comment1.CreateDate.Format("01-02")
		temp := &comment.InternalComment{
			Id:         int32(comment1.ID),
			UserId:     int32(comment1.UserID),
			Content:    comment1.Content,
			CreateDate: createDate,
		}
		res = append(res, temp)
	}
	resp.CommentList = res
	return resp, nil
}

// InternalCountComment implements the InternalCommentServiceImpl interface.
func (s *InternalCommentServiceImpl) InternalCountComment(ctx context.Context, request *comment.InternalCountCommentRequest) (resp *comment.InternalCountCommentResponse, err error) {
	// TODO: Your code here...

	return
}
