package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
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
func GetAllPosts(ctx context.Context, page int, query, category string) ([]Post, error) {
	var posts []Post
	offset := (page - 1) * PAGE_SIZE
	dbQuery := conn.WithContext(ctx).Order("created_at desc").Offset(offset).Limit(PAGE_SIZE)

	if query != "" {
		queryMatcher := fmt.Sprintf("%%%s%%", query)
		markupMatcher := fmt.Sprintf("%% %s %%", query)
		dbQuery = dbQuery.Where("title LIKE ? OR category LIKE ? OR markup LIKE ?", queryMatcher, queryMatcher, markupMatcher)
	}

	if category != "" {
		categoryMatcher := fmt.Sprintf("%%%s%%", category)
		dbQuery = dbQuery.Where("category LIKE ?", categoryMatcher)
	}

	result := dbQuery.Find(&posts)
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

func GetPostById(ctx context.Context, id int) (*Post, error) {
	var post Post
	result := conn.WithContext(ctx).First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("get post by id error: %w", ErrNotFound)
		}
		return nil, fmt.Errorf("get post by id error: %w", result.Error)
	}
	return &post, nil
}

func DropPosts(ctx context.Context) error {
	return conn.WithContext(ctx).Migrator().DropTable(&Post{})
}

func ReloadPosts(ctx context.Context) error {
	outputDirRead, _ := os.Open("./posts")

	postFiles, err := outputDirRead.ReadDir(0)
	if err != nil {
		return err
	}

	for _, postFile := range postFiles {
		// Get name of file.
		postFileName := postFile.Name()

		// read markdown contents
		md, err := os.ReadFile("./posts/" + postFileName)
		if err != nil {
			return err
		}

		// parse markdown file into html
		parser := parser.New()
		html := markdown.ToHTML(md, parser, nil)

		// parse filename into post metadata
		parts := strings.Split(postFileName, "$")

		date, err := time.Parse("2006-01-02", parts[2])
		if err != nil {
			return err
		}

		conn.WithContext(ctx).Create(&Post{
			Title:     parts[0],
			Markup:    string(html),
			Category:  parts[1],
			Upvotes:   0,
			CreatedAt: date,
		})
	}
	return nil
}

func UpvotePost(ctx context.Context, id int) (int, error) {
	result := conn.WithContext(ctx).Exec("UPDATE posts SET upvotes = upvotes + 1 WHERE id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("upvote post by id error: %w", ErrNotFound)
		}
		return 0, fmt.Errorf("upvote post by id error: %w", result.Error)
	}
	var post Post
	result = conn.WithContext(ctx).First(&post, id)
	if result.Error != nil {
		return 0, fmt.Errorf("get post by id error: %w", result.Error)
	}
	return int(post.Upvotes), nil
}
