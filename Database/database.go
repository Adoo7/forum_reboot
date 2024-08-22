package database

import (
	"database/sql"
	"errors"
	"fmt"
	"forum_reboot/structs"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Connect connects to the database file specified by dbPath
func Connect(dbPath string) error {
	// sql.Open won't error if file not found
	fi, err := os.Stat(dbPath)
	if err != nil || fi.IsDir() {
		return errors.New("database file not found")
	}
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", dbPath)
	ldb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("can't connect to database: %w", err)
	}
	DB = ldb
	return nil
}

// Close closes the database connection
func Close() error {
	return DB.Close()
}

// CheckExistance checks if a value exists in a specific table and column
func CheckExistance(tablename, columnname, value string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", tablename, columnname)
	stmt, err := DB.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(value).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// InsertCategory inserts a new category into the Category table
func InsertCategory(category structs.CategoryResponse) error {
	query := `INSERT INTO Category (Name, Description) VALUES (?, ?)`
	_, err := DB.Exec(query, category.Name, category.Description)
	return err
}

// GetCategory retrieves a category by its ID
func GetCategory(id int) (structs.CategoryResponse, error) {
	var category structs.CategoryResponse
	query := `SELECT Name, Description FROM Category WHERE Category_ID = ?`
	row := DB.QueryRow(query, id)
	err := row.Scan(&category.Name, &category.Description)
	if err != nil {
		return category, err
	}
	return category, nil
}

// InsertUser inserts a new user into the User table
func InsertUser(user structs.UserResponse, hashedPassword string) error {
	query := `INSERT INTO User (username, email, passwords) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, user.Username, user.Email, hashedPassword)
	return err
}

// GetUser retrieves a user by their username
func GetUser(username string) (structs.UserResponse, error) {
	var user structs.UserResponse
	query := `SELECT username, email, passwords FROM User WHERE username = ?`
	row := DB.QueryRow(query, username)
	err := row.Scan(&user.Username, &user.Email, &user.Passwords)
	if err != nil {
		return user, err
	}
	return user, nil
}

// InsertPost inserts a new post into the Post table
func InsertPost(post structs.PostResponse) error {
	query := `INSERT INTO Post (User_ID, Title, Messages, Like_count, DisLike_Count) VALUES (?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, post.User_ID, post.Title, post.Message, post.Like_count, post.DisLike_Count)
	return err
}

// GetPost retrieves a post by its ID
func GetPost(id int) (structs.PostResponse, error) {
	var post structs.PostResponse
	query := `SELECT User_ID, Title, Messages, Like_count, DisLike_Count FROM Post WHERE PostID = ?`
	row := DB.QueryRow(query, id)
	err := row.Scan(&post.User_ID, &post.Title, &post.Message, &post.Like_count, &post.DisLike_Count)
	if err != nil {
		return post, err
	}
	return post, nil
}

// InsertComment inserts a new comment into the Comment table
func InsertComment(comment structs.CommentResponse) error {
	query := `INSERT INTO Comment (User_ID, PostID, message, Like_Count, DisLike_Count) VALUES (?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, comment.User_ID, comment.PostID, comment.Message, comment.Like_Count, comment.DisLike_Count)
	return err
}

// GetComment retrieves a comment by its ID
func GetComment(id int) (structs.CommentResponse, error) {
	var comment structs.CommentResponse
	query := `SELECT User_ID, PostID, message, Like_Count, DisLike_Count FROM Comment WHERE Comment_ID = ?`
	row := DB.QueryRow(query, id)
	err := row.Scan(&comment.User_ID, &comment.PostID, &comment.Message, &comment.Like_Count, &comment.DisLike_Count)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

// InsertReaction inserts a new reaction into the Reaction table
func InsertReaction(reaction structs.ReactionResponse) error {
	query := `INSERT INTO Reaction (User_ID, PostID, Comment_ID, Type) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, reaction.User_ID, reaction.PostID, reaction.Comment_ID, reaction.Type)
	return err
}

// GetReaction retrieves a reaction by its ID
func GetReaction(id int) (structs.ReactionResponse, error) {
	var reaction structs.ReactionResponse
	query := `SELECT User_ID, PostID, Comment_ID, Type FROM Reaction WHERE ReactionID = ?`
	row := DB.QueryRow(query, id)
	err := row.Scan(&reaction.User_ID, &reaction.PostID, &reaction.Comment_ID, &reaction.Type)
	if err != nil {
		return reaction, err
	}
	return reaction, nil
}
