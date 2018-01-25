// Package models contains wordpress data models
package models

// Post is wordpress data structure.
// Title - the title of the given post
// Content - the content of the given post
type Post struct {
	Title string `json:"title"`
	Content  string `json:"content"`
}