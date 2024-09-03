package server

import (
    "encoding/json"
    "net/http"
    "fmt"
)

func updatePostLikeHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request to update like/dislike")

    var data struct {
        PostID int    `json:"post_id"`
        Action string `json:"action"`
    }

    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        log.Println("Error decoding request body:", err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    log.Printf("Updating post %d with action %s", data.PostID, data.Action)

    // Update database (replace this with actual database logic)
    success := updateLikeDislikeInDB(data.PostID, data.Action)

    if !success {
        log.Println("Failed to update post in database")
        http.Error(w, "Failed to update post", http.StatusInternalServerError)
        return
    }

    log.Println("Post updated successfully")
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func updateLikeDislikeInDB(postID int, action string) bool {
    // Logic to update like/dislike in your database
    // Return true if the update was successful, otherwise false
    return true
}


