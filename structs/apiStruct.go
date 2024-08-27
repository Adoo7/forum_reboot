package structs

// MARK: Post API response
type PostAPIResponse struct {
	PostID       	int                    	`json:"post_id"`
	UserID      	int                    	`json:"user_id"`
	Title        	string                 	`json:"title"`
	Message      	string                 	`json:"message"`
	TimeofPost   	string                 	`json:"time_of_post"`
	LikeCount    	int                    	`json:"like_count"`
	DislikeCount 	int                    	`json:"dislike_count"`
	Categories   	*[]CategoryAPIResponse 	`json:"categories"`
	Comments			*[]CommentAPIResponse 	`json:"comments"`
}

type CategoryAPIResponse struct {
	CategoryID  	int    	`json:"category_id"`
	Name        	string 	`json:"name"`
	Description 	string 	`json:"description"`
}

type CommentAPIResponse struct {
	CommentID 		int 		`json:"comment_id"`
	UserID 				int 		`json:"user_id"`
	PostID        int    	`json:"post_id"`
	Message 			string 	`json:"message"`
	TimeofComment string 	`json:"time_of_post"`
	LikeCount 		int 		`json:"like_count"`
	DislikeCount 	int 		`json:"dislike_count"`
}