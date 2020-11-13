package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents an user´s post
type Post struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	UserID    uint64    `json:"userID,omitempty"`
	UserNick  string    `json:"userNick,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare validates and format the post
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("Invalid blank title")
	}

	if post.Content == "" {
		return errors.New("Invalid blank content")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
