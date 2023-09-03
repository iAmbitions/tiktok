namespace go douyin.favorite

include "feed.thrift"

struct InternalFavoriteRequest {
  1: i32 actor_id,     // 用户id
  2: i32 video_id,     // 视频id
  3: i32 action_type   // 1-点赞，2-取消点赞
}

struct InternalFavoriteResponse {
  1: i32 status_code,  // 状态码，0-成功，其他值-失败
  2: optional string status_msg   // 返回状态描述
}

struct InternalFavoriteListRequest {
  1: i32 actor_id,     // 发出请求的用户的id
  2: i32 user_id       // 用户id
}

struct InternalFavoriteListResponse {
  1: i32 status_code,  // 状态码，0-成功，其他值-失败
  2: optional string status_msg,  // 返回状态描述
  3: list<feed.Video> video_list  // 用户点赞视频列表
}

struct InternalIsFavoriteRequest {
  1: i32 user_id,      // 用户id
  2: i32 video_id      // 视频id
}

struct InternalIsFavoriteResponse {
  1: bool result       // 结果
}

struct InternalCountFavoriteRequest {
  1: i32 video_id      // 视频id
}

struct InternalCountFavoriteResponse {
  1: i32 status_code,
  2: optional string status_msg,
  3: i32 count         // 点赞数
}

struct InternalCountUserFavoriteRequest {
  1: i32 user_id       // 用户id
}

struct InternalCountUserFavoriteResponse {
  1: i32 status_code,
  2: optional string status_msg,
  3: i32 count         // 点赞数
}

struct InternalCountUserTotalFavoritedRequest {
  1: i32 actor_id,
  2: i32 user_id
}

struct InternalCountUserTotalFavoritedResponse {
  1: i32 status_code,
  2: optional string status_msg,
  3: i32 count         // 点赞数
}

service InternalFavoriteService {
  InternalFavoriteResponse InternalFavoriteAction(1: InternalFavoriteRequest request),
  InternalFavoriteListResponse InternalFavoriteList(1: InternalFavoriteListRequest request),
  InternalIsFavoriteResponse InternalIsFavorite(1: InternalIsFavoriteRequest request),
  InternalCountFavoriteResponse InternalCountFavorite(1: InternalCountFavoriteRequest request),
  InternalCountUserFavoriteResponse InternalCountUserFavorite(1: InternalCountUserFavoriteRequest request),
  InternalCountUserTotalFavoritedResponse InternalCountUserTotalFavorited(1: InternalCountUserTotalFavoritedRequest request)
}
