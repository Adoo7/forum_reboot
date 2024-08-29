package server

import (
	"encoding/json"
	"forum_reboot/structs"
	"net/http"
	"time"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	userID, valid := CheckSession(r)
	if !valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	
	log.Printf("User ID from session: %d", userID)
	
	var newComment structs.CommentAPIResponse
	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		http.Error(w, "ERR:400 Invalid input", http.StatusBadRequest)
		return
	}
	
	newComment.TimeofComment = time.Now().Format("2006-01-02 15:04:05")
	newComment.UserID = userID
	newComment.LikeCount = 0
	newComment.DislikeCount = 0
	
	_, err = DB.Exec("INSERT INTO Comment (User_ID, PostID, TimeofComment, message, Like_Count, DisLike_Count) VALUES (?, ?, ?, ?, ?, ?)",
		newComment.UserID, newComment.PostID, newComment.TimeofComment, newComment.Message, newComment.LikeCount, newComment.DislikeCount)
	if err != nil {
		http.Error(w, "ERR:401 Failed to create comment", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Comment created successfully",
	})
}