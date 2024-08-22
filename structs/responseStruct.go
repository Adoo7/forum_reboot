package structs

// MARK: Category
type CategoryResponse struct {
    CategoryID   int    `json:"category_id"`
    Name         string `json:"name"`
    Description  string `json:"description"`
}

// MARK: User
type UserResponse struct {
    UserID    int    `json:"user_id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    CreatedAt string `json:"created_at"`
}

// MARK: UserRole
type UserRole struct {
    UserRoleID int    `json:"user_role_id"`
    UserID     int    `json:"user_id"`
    RoleName   string `json:"role_name"`
    CanPost    bool   `json:"can_post"`
    CanLike    bool   `json:"can_like"`
    CanComment bool   `json:"can_comment"`
    CanDislike bool   `json:"can_dislike"`
}

// MARK: UserSession
type UserSession struct {
    UserSessionID string `json:"user_session_id"`
    UserID        int    `json:"user_id"`
    Token         string `json:"token"`
    CreationTime  string `json:"creation_time"` // Use a suitable time format for JSON
    ExpireTime    string `json:"expire_time"`   // Use a suitable time format for JSON
}

// MARK: Post
type PostResponse struct {
    PostID        int    `json:"post_id"`
    UserID        int    `json:"user_id"`
    Title         string `json:"title"`
    Message       string `json:"message"`
    TimeofPost    string `json:"time_of_post"`
    LikeCount     int    `json:"like_count"`
    DislikeCount  int    `json:"dislike_count"`
    // Removed Categories field
}

// MARK: PostCategory
type PostCategory struct {
    PostCategoryID int `json:"post_category_id"`
    CategoryID     int `json:"category_id"`
    PostID         int `json:"post_id"`
}

// MARK: Comment
type CommentResponse struct {
    CommentID      int    `json:"comment_id"`
    UserID         int    `json:"user_id"`
    PostID         int    `json:"post_id"`
    TimeofComment  string `json:"time_of_comment"`
    Message        string `json:"message"`
    LikeCount      int    `json:"like_count"`
    DislikeCount   int    `json:"dislike_count"`
}

// MARK: Reaction
type ReactionResponse struct {
    ReactionID     int    `json:"reaction_id"`
    UserID         int    `json:"user_id"`
    PostID         *int   `json:"post_id"`
    CommentID      *int   `json:"comment_id"`
    Type           string `json:"type"`
    TimeofReaction string `json:"time_of_reaction"`
}
