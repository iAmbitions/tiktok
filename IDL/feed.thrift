namespace go douyin.feed

include "user.thrift"

struct Video {
  1: i32 id,
  2: user.User author,
  3: string play_url,
  4: string cover_url,
  5: i32 favorite_count,
  6: i32 comment_count,
  7: bool is_favorite,
  8: string title
}

struct ListFeedRequest {
  1: optional string latest_time, // 限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
  2: optional i32 actor_id // 发送请求的用户的id
}

struct ListFeedResponse {
  1: i32 status_code,
  2: optional string status_msg,
  3: optional i64 next_time,
  4: list<Video> video_list
}

struct QueryVideosRequest {
  1: i32 actor_id,
  2: list<i32> video_ids
}

struct QueryVideosResponse {
  1: i32 status_code,
  2: optional string status_msg,
  3: list<Video> video_list
}

service FeedService {
  ListFeedResponse ListVideos(1: ListFeedRequest request),
  QueryVideosResponse QueryVideos(1: QueryVideosRequest request)
}
