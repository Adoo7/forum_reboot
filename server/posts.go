package server

import (
	"encoding/json"
	"time"
	"strconv"

	"forum_reboot/structs"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {

	// Get all posts
	postRows, err := DB.Query("SELECT PostID, User_ID, Title, Messages, TimeofPost, Like_count, DisLike_Count FROM Post")
	if err != nil {
		println(err.Error())
		http.Error(w, "ERR:100 Failed to get posts", http.StatusInternalServerError)
		return
	}
	defer postRows.Close()

	posts := []structs.PostAPIResponse{}

	for postRows.Next() {
		i := structs.PostAPIResponse{}
		err = postRows.Scan(&i.PostID, &i.UserID, &i.Title, &i.Message, &i.TimeofPost, &i.LikeCount, &i.DislikeCount)
		if err != nil {
			println(err.Error())
			http.Error(w, "ERR:101 Failed to get posts", http.StatusInternalServerError)
			return
		}

		// Initialize Categories slice to avoid nil pointer dereference
		i.Categories = &[]structs.CategoryAPIResponse{}

		println("postid:", i.PostID)

		// Get all categories for the post
		catRows, err := DB.Query("SELECT Category.Category_ID, Category.Name, Category.Description FROM Category RIGHT JOIN PostCategory ON Category.Category_ID = PostCategory.Category_ID WHERE PostCategory.PostID = ?", i.PostID)
		if err != nil {
			println(err.Error())
			http.Error(w, "ERR:102 Failed to get categories for post", http.StatusInternalServerError)
			return
		}
		defer catRows.Close()

		for catRows.Next() {
			c := structs.CategoryAPIResponse{}
			err = catRows.Scan(&c.CategoryID, &c.Name, &c.Description)
			if err != nil {
				println(err.Error())
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
		println(err.Error())
		http.Error(w, "ERR:104 Failed to encode posts", http.StatusInternalServerError)
		return
	}
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
    userID, valid := CheckSession(r)
    if !valid {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var newPost structs.PostAPIResponse
    err := json.NewDecoder(r.Body).Decode(&newPost)
    if err != nil {
        http.Error(w, "ERR:200 Invalid input", http.StatusBadRequest)
        return
    }

    newPost.TimeofPost = time.Now().Format("2006-01-02 15:04:05")
    newPost.UserID = userID
    println("Creating post with UserID:", newPost.UserID) // Debug log

    res, err := DB.Exec("INSERT INTO Post (User_ID, Title, Messages, TimeofPost, Like_count, DisLike_Count) VALUES (?, ?, ?, ?, ?, ?)",
        newPost.UserID, newPost.Title, newPost.Message, newPost.TimeofPost, newPost.LikeCount, newPost.DislikeCount)
    if err != nil {
        http.Error(w, "ERR:201 Failed to create post", http.StatusInternalServerError)
        return
    }

    postID, err := res.LastInsertId()
    if err != nil {
        http.Error(w, "ERR:202 Failed to retrieve post ID", http.StatusInternalServerError)
        return
    }

    if newPost.Categories != nil {
        for _, category := range *newPost.Categories {
            _, err := DB.Exec("INSERT INTO PostCategory (PostID, Category_ID) VALUES (?, ?)", postID, category.CategoryID)
            if err != nil {
                http.Error(w, "ERR:203 Failed to associate categories with post", http.StatusInternalServerError)
                return
            }
        }
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "post_id": postID,
    })
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	// Get the post ID from the query parameters
	postIDParam := r.URL.Query().Get("id")
	if postIDParam == "" {
		http.Error(w, "ERR:300 Post ID is required", http.StatusBadRequest)
		return
	}

	// Convert the post ID to an integer
	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		println(err.Error())
		http.Error(w, "ERR:301 Invalid post ID", http.StatusBadRequest)
		return
	}

	// Get the post details
	post := structs.PostAPIResponse{}
	err = DB.QueryRow("SELECT PostID, User_ID, Title, Messages, TimeofPost, Like_count, DisLike_Count FROM Post WHERE PostID = ?", postID).
		Scan(&post.PostID, &post.UserID, &post.Title, &post.Message, &post.TimeofPost, &post.LikeCount, &post.DislikeCount)
	if err != nil {
		println(err.Error())
		http.Error(w, "ERR:302 Failed to get post details", http.StatusInternalServerError)
		return
	}

	// Initialize Categories slice to avoid nil pointer dereference
	post.Categories = &[]structs.CategoryAPIResponse{}

	// Get all categories for the post
	catRows, err := DB.Query("SELECT Category.Category_ID, Category.Name, Category.Description FROM Category RIGHT JOIN PostCategory ON Category.Category_ID = PostCategory.Category_ID WHERE PostCategory.PostID = ?", post.PostID)
	if err != nil {
		println(err.Error())
		http.Error(w, "ERR:303 Failed to get categories for post", http.StatusInternalServerError)
		return
	}
	defer catRows.Close()

	for catRows.Next() {
		c := structs.CategoryAPIResponse{}
		err = catRows.Scan(&c.CategoryID, &c.Name, &c.Description)
		if err != nil {
			println(err.Error())
			http.Error(w, "ERR:304 Failed to get categories for post", http.StatusInternalServerError)
			return
		}
		*post.Categories = append(*post.Categories, c)
	}

	// Initialize Comments slice to avoid nil pointer dereference
	post.Comments = &[]structs.CommentAPIResponse{}

	// Get all comments for the post
	commentRows, err := DB.Query("SELECT Comment_ID, User_ID, message, TimeofComment, Like_Count, Dislike_Count FROM Comment WHERE PostID = ?", post.PostID)
	if err != nil {
		println(err.Error())
		http.Error(w, "ERR:305 Failed to get comments for post", http.StatusInternalServerError)
		return
	}
	defer commentRows.Close()

	for commentRows.Next() {
		c := structs.CommentAPIResponse{}
		err = commentRows.Scan(&c.CommentID, &c.UserID, &c.Message, &c.TimeofComment, &c.LikeCount, &c.DislikeCount)
		if err != nil {
			println(err.Error())
			http.Error(w, "ERR:306 Failed to get comments for post", http.StatusInternalServerError)
			return
		}
		*post.Comments = append(*post.Comments, c)
	}

	// Respond with the post details, categories, and comments
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		println(err.Error())
		http.Error(w, "ERR:307 Failed to encode post details", http.StatusInternalServerError)
		return
	}
}