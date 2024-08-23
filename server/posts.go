package server

import (
	"encoding/json"
	"net/http"
	"forum_reboot/structs"
_ "github.com/mattn/go-sqlite3" 
)


func GetPosts(w http.ResponseWriter, r *http.Request) {
        
	rows, err := DB.Query("SELECT * FROM Post")
	if err != nil {
			http.Error(w, "Failed to get posts", http.StatusBadRequest)
			return
	}
	
	defer rows.Close()
	data := []structs.PostResponse{}

	for rows.Next() {
			i := structs.PostResponse{}
			err = rows.Scan(&i.PostID,&i.UserID,&i.Title,&i.Message,&i.TimeofPost,&i.LikeCount,&i.DislikeCount)
			if err != nil {
					http.Error(w, "Failed to get posts", http.StatusInternalServerError)
					return
			}
			data = append(data, i)
	}
	
	w.Header().Set("Content-Type", "application/json")
	d, err := json.Marshal(data) 
	if err != nil {
			http.Error(w, "Failed to get posts", http.StatusInternalServerError)
			return
	}

	w.Write([]byte(d))
}