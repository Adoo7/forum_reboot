package structs

import "time"

// UserSessionDB represents a user session in the database
type UserSessionDB struct {
	Id           string    `json:"id"`
	UserId       int       `json:"user_id"`
	Token        string    `json:"token"`
	CreationTime time.Time `json:"creation_time"`
	ExpireTime   time.Time `json:"expire_time"`
}

// UserSessionAPI represents a user session in API responses
type UserSessionAPI struct {
	UserSessionID string `json:"user_session_id"`
	UserID        int    `json:"user_id"`
	Token         string `json:"token"`
	CreationTime  string `json:"creation_time"`
	ExpireTime    string `json:"expire_time"`
}

// Category represents a category for posts
type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UserRole represents the role assigned to a user
type UserRole struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	RoleName   string `json:"role_name"`
	CanPost    bool   `json:"can_post"`
	CanLike    bool   `json:"can_like"`
	CanComment bool   `json:"can_comment"`
	CanDislike bool   `json:"can_dislike"`
}

// Post represents a forum post
type Post struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`
	Title         string    `json:"title"`
	Message       string    `json:"message"`
	TimeofPost    time.Time `json:"timeof_post"`
	LikeCount     int       `json:"like_count"`
	DislikeCount  int       `json:"dislike_count"`
}

// PostCategory represents the relationship between posts and categories
type PostCategory struct {
	Id         int `json:"id"`
	PostId     int `json:"post_id"`
	CategoryId int `json:"category_id"`
}

// Comment represents a comment on a post
type Comment struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`
	PostId        int       `json:"post_id"`
	TimeofComment time.Time `json:"timeof_comment"`
	Message       string    `json:"message"`
	LikeCount     int       `json:"like_count"`
	DislikeCount  int       `json:"dislike_count"`
}

// Reaction represents a reaction to either a post or comment
type Reaction struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	PostId         *int      `json:"post_id"`     // Pointer to handle possible NULL values
	CommentId      *int      `json:"comment_id"`  // Pointer to handle possible NULL values
	Type           string    `json:"type"`
	TimeofReaction time.Time `json:"timeof_reaction"`
}
