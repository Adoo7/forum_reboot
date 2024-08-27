package server

import (
	"encoding/json"
	"forum_reboot/structs"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var newComment structs.CommentAPIResponse

	// Decode the request body into the newComment struct
	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		http.Error(w, "ERR:400 Invalid input", http.StatusBadRequest)
		return
	}

	// Set the TimeofComment to the current time
	newComment.TimeofComment = time.Now().Format("2006-01-02 15:04:05")

	// Set default values for Like_Count and DisLike_Count
	newComment.LikeCount = 0
	newComment.DislikeCount = 0

	// Insert the new comment into the Comment table
	_, err = DB.Exec("INSERT INTO Comment (User_ID, PostID, TimeofComment, message, Like_Count, DisLike_Count) VALUES (?, ?, ?, ?, ?, ?)",
		newComment.UserID, newComment.PostID, newComment.TimeofComment, newComment.Message, newComment.LikeCount, newComment.DislikeCount)
	if err != nil {
		http.Error(w, "ERR:401 Failed to create comment", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Comment created successfully",
	})
}
