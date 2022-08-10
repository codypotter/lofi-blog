package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const PAGE_SIZE = 10

// Add a post to the db
func AddPost(title string, markup string, category string, date time.Time) {
	conn.Create(&Post{
		Title:     title,
		Markup:    markup,
		Category:  category,
		Upvotes:   0,
		CreatedAt: date,
	})
}

// Get all posts from db with pagination
func GetAllPosts(ctx context.Context, page int) ([]Post, error) {
	var posts []Post
	offset := (page - 1) * PAGE_SIZE
	result := conn.WithContext(ctx).Order("created_at desc").Offset(offset).Limit(PAGE_SIZE).Find(&posts)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []Post{}, nil
		}
		return nil, fmt.Errorf("get all posts error: %w", ErrConnection)
	}
	return posts, nil
}

// Get most recently created post from db
func GetMostRecentPost(ctx context.Context) (*Post, error) {
	var featuredPost Post
	result := conn.WithContext(ctx).Order("created_at desc").Limit(1).Find(&featuredPost)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("get featured post error: %w", ErrNotFound)
		}
		return nil, fmt.Errorf("get featured post error: %w", result.Error)
	}
	return &featuredPost, nil
}
