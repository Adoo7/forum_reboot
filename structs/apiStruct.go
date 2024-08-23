package structs 

// MARK: Post API response
type PostAPIResponse struct {
	PostID        int    `json:"post_id"`
	UserID        int    `json:"user_id"`
	Title         string `json:"title"`
	Message       string `json:"message"`
	TimeofPost    string `json:"time_of_post"`
	LikeCount     int    `json:"like_count"`
	DislikeCount  int    `json:"dislike_count"`
	Categories 		*[]CategoryAPIResponse `json:"categories"`
}

type CategoryAPIResponse struct {
	CategoryID int `json:"category_id"`
	Name string `json:"name"`
	Description string `json:"description"`
}