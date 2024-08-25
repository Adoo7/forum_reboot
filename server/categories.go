package server

import (
	"encoding/json"

	"forum_reboot/structs"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {

    // Get all categories
    catRows, err := DB.Query("SELECT * FROM Category")
    if err != nil {
        http.Error(w, "ERR:100 Failed to get categories", http.StatusInternalServerError)
        return
    }
    defer catRows.Close()

    cats := []structs.CategoryAPIResponse{}

    for catRows.Next() {
        i := structs.CategoryAPIResponse{}
        err = catRows.Scan(&i.CategoryID, &i.Name, &i.Description)
        if err != nil {
            http.Error(w, "ERR:101 Failed to get categories", http.StatusInternalServerError)
            return
        }

        cats = append(cats, i)
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(cats)
    if err != nil {
        http.Error(w, "ERR:104 Failed to encode categories", http.StatusInternalServerError)
        return
    }
}
