package server

import (
	"encoding/json"
	"time"

	"forum_reboot/structs"
	"net/http"

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

		println("postid:", i.PostID)

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
				println(err.Error())
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost structs.PostAPIResponse

	// Decode the request body into the newPost struct
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "ERR:200 Invalid input", http.StatusBadRequest)
		return
	}

	// Set the TimeofPost to the current time
	newPost.TimeofPost = time.Now().Format("2006-01-02 15:04:05")
	println("userID", newPost.UserID)
	// Insert the new post into the Post table
	res, err := DB.Exec("INSERT INTO Post (User_ID, Title, Messages, TimeofPost, Like_count, DisLike_Count) VALUES (?, ?, ?, ?, ?, ?)",
		newPost.UserID, newPost.Title, newPost.Message, newPost.TimeofPost, newPost.LikeCount, newPost.DislikeCount)
	if err != nil {
		println(err.Error())
		http.Error(w, "ERR:201 Failed to create post", http.StatusInternalServerError)
		return
	}

	// Get the last inserted PostID
	postID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "ERR:202 Failed to retrieve post ID", http.StatusInternalServerError)
		return
	}

	// Add categories if provided
	if newPost.Categories != nil {
		for _, category := range *newPost.Categories {
			_, err := DB.Exec("INSERT INTO PostCategory (PostID, Category_ID) VALUES (?, ?)", postID, category.CategoryID)
			if err != nil {
				http.Error(w, "ERR:203 Failed to associate categories with post", http.StatusInternalServerError)
				return
			}
		}
	}

	// Respond with the created post ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"post_id": postID,
	})
}
