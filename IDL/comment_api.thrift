namespace go douyin.comment
include "user.thrift"
struct User {
  1: i32 id,      // 用户id
  2: string name, // 用户名称
  3: i32 follow_count,  // 关注总数
  4: i32 follower_count,  // 粉丝总数
  5: bool is_follow,   // true-已关注，false-未关注
  6: optional string avatar,  // 用户头像
  7: optional string background_image,  // 用户个人顶部大图
  8: optional string signature,   // 个人简介
  9: optional i32 total_favorited,  // 获赞数量
 10: optional i32 work_count,   // 作品数量
 11: optional i32 favorite_count   // 点赞数量
}
struct Comment {
  1: i32 id
  2: i32 user_id
  3: string content
  4: string create_date
}


enum ActionCommentType {
  ACTION_COMMENT_TYPE_UNSPECIFIED = 0
  ACTION_COMMENT_TYPE_ADD = 1
  ACTION_COMMENT_TYPE_DELETE = 2
}


struct ActionCommentRequest {
    1: string token,
    2: i32 video_id,
    3: ActionCommentType action_type,
    4: optional string comment_text,
    5: optional i32 comment_id
}

struct ActionCommentResponse {
  1: i32 status_code
  2: string status_msg
  3: Comment comment
}

struct ListCommentRequest {
      1: string token
  2: i32 video_id
}

struct ListCommentResponse {
  1: i32 status_code
  2: string status_msg
  3: list<Comment> comment_list
}

struct CountCommentRequest {
    1: i32 actor_id,
    2: i32 video_id
}

struct CountCommentResponse {
    1: i32 status_code,
    2: optional string status_msg,
    3: i32 comment_count
}

service CommentService {
    ActionCommentResponse ActionComment(1: ActionCommentRequest request),
    ListCommentResponse ListComment(1: ListCommentRequest request),
    CountCommentResponse CountComment(1: CountCommentRequest request)
}