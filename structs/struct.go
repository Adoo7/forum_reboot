
package structs

import "time"

// User represents a user in the forum
type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

// Category represents a category for posts
type Category struct {
	Id          int
	Name        string
	Description string
}

// UserRole represents the role assigned to a user
type UserRole struct {
	Id         int
	UserId     int
	RoleName   string
	CanPost    bool
	CanLike    bool
	CanComment bool
	CanDislike bool
}

// UserSession represents a user session for authentication
type UserSession struct {
	Id           string
	UserId       int
	Token        string
	CreationTime time.Time
	ExpireTime   time.Time
}

// Post represents a forum post
type Post struct {
	Id            int
	UserId        int
	Title         string
	Message       string
	TimeofPost    time.Time
	LikeCount     int
	DislikeCount  int
}

// PostCategory represents the relationship between posts and categories
type PostCategory struct {
	Id         int
	PostId     int
	CategoryId int
}

// Comment represents a comment on a post
type Comment struct {
	Id            int
	UserId        int
	PostId        int
	TimeofComment time.Time
	Message       string
	LikeCount     int
	DislikeCount  int
}

// Reaction represents a reaction to either a post or comment
type Reaction struct {
	Id             int
	UserId         int
	PostId         *int // Pointer to allow NULL values
	CommentId      *int // Pointer to allow NULL values
	Type           string
	TimeofReaction time.Time
}
