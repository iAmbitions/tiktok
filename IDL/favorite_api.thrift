namespace go douyin.api.favorite

include "feed.thrift"

struct FavoriteRequest {
      1: string token
  2: i32 video_id,     // 视频id
  3: i32 action_type   // 1-点赞，2-取消点赞
}

struct FavoriteResponse {
  1: i32 status_code,  // 状态码，0-成功，其他值-失败
  2: optional string status_msg   // 返回状态描述
}

struct FavoriteListRequest {
      1: string token
  2: i32 user_id       // 用户id
}

struct FavoriteListResponse {
  1: i32 status_code,  // 状态码，0-成功，其他值-失败
  2: optional string status_msg,  // 返回状态描述
  3: list<feed.Video> video_list  // 用户点赞视频列表
}

struct IsFavoriteRequest {
      1: string token
  2: i32 video_id      // 视频id
}

struct IsFavoriteResponse {
  1: bool result       // 结果
}

struct CountFavoriteRequest {
  1: i32 video_id      // 视频id
}

struct CountFavoriteResponse {
  1: i32 status_code,
  2: optional string status_msg,
  3: i32 count         // 点赞数
}

struct CountUserFavoriteRequest {
      1: string token
}

struct CountUserFavoriteResponse {
  1: i32 status_code,
  2: optional string status_msg,
  3: i32 count         // 点赞数
}

struct CountUserTotalFavoritedRequest {
  1: i32 actor_id,
      2: string token
}

struct CountUserTotalFavoritedResponse {
  1: i32 status_code,
  2: optional string status_msg,
  3: i32 count         // 点赞数
}

service FavoriteService {
  FavoriteResponse FavoriteAction(1: FavoriteRequest request),
  FavoriteListResponse FavoriteList(1: FavoriteListRequest request),
  IsFavoriteResponse IsFavorite(1: IsFavoriteRequest request),
  CountFavoriteResponse CountFavorite(1: CountFavoriteRequest request),
  CountUserFavoriteResponse CountUserFavorite(1: CountUserFavoriteRequest request),
  CountUserTotalFavoritedResponse CountUserTotalFavorited(1: CountUserTotalFavoritedRequest request)
}
