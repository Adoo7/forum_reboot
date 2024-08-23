package server

import (
	"encoding/json"
	"net/http"
	"forum_reboot/structs"
_ "github.com/mattn/go-sqlite3" 
)


func GetPosts(w http.ResponseWriter, r *http.Request) {

	// Get all posts
	postRows, err := DB.Query("SELECT PostID, User_ID, Title, Messages, TimeofPost, Like_count, DisLike_Count FROM Post")
	if err != nil {
		http.Error(w, "ERR:100 Failed to get posts", http.StatusInternalServerError)
		return
	}
	defer postRows.Close()

	posts := []structs.PostAPIResponse{}

	for postRows.Next() {
		i := structs.PostAPIResponse{}
		err = postRows.Scan(&i.PostID, &i.UserID, &i.Title, &i.Message, &i.TimeofPost, &i.LikeCount, &i.DislikeCount)
		if err != nil {
			http.Error(w, "ERR:101 Failed to get posts", http.StatusInternalServerError)
			return
		}

		// Initialize Categories slice to avoid nil pointer dereference
		i.Categories = &[]structs.CategoryAPIResponse{}

		// Get all categories for the post
		catRows, err := DB.Query("SELECT Category.Category_ID, Category.Name, Category.Description FROM Category RIGHT JOIN PostCategory ON Category.Category_ID = PostCategory.Category_ID WHERE PostCategory.PostID = ?", i.PostID)
		if err != nil {
			http.Error(w, "ERR:102 Failed to get categories for post", http.StatusInternalServerError)
			return
		}
		defer catRows.Close()

		for catRows.Next() {
			c := structs.CategoryAPIResponse{}
			err = catRows.Scan(&c.CategoryID, &c.Name, &c.Description)
			if err != nil {
				http.Error(w, "ERR:103 Failed to get categories for post", http.StatusInternalServerError)
				return
			}
			*i.Categories = append(*i.Categories, c)
		}

		posts = append(posts, i)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, "ERR:104 Failed to encode posts", http.StatusInternalServerError)
		return
	}
}