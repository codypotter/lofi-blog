package db

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var conn *gorm.DB

type Post struct {
	ID        uint
	Title     string
	Markup    string
	Category  string
	Upvotes   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Connect() {
	var err error
	conn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
	}

	if conn.AutoMigrate(&Post{}) != nil {
		log.Panic("Failed to automigrate database")
	}

	readPosts()
}

func AddPost(title string, markup string, category string, date time.Time) {
	log.Println("adding post...")
	conn.Create(&Post{
		Title:    title,
		Markup:   markup,
		Category: category,
		Upvotes:  0,
	})
}

func readPosts() {
	outputDirRead, _ := os.Open("./posts")

	postFiles, err := outputDirRead.ReadDir(0)
	if err != nil {
		log.Fatal("Failed to read posts directory")
	}

	// Iterate over files in posts directory, creating a post for each
	for _, postFile := range postFiles {
		// Get name of file.
		postFileName := postFile.Name()

		// Print name.
		log.Println(postFileName)

		// read markdown contents
		md, err := os.ReadFile("./posts/" + postFileName)
		if err != nil {
			log.Fatalf("Error reading file %s\n", postFileName)
		}

		// parse markdown file into html
		parser := parser.New()
		html := markdown.ToHTML(md, parser, nil)

		// parse filename into post metadata
		title, category, date := parseFileName(postFileName)
		fmt.Printf("Title: %s, html: %s, category: %s, date: %v", title, string(html), category, date)

		AddPost(title, string(html), category, date)
	}
}

func parseFileName(filename string) (string, string, time.Time) {
	parts := strings.Split(filename, "$")

	date, err := time.Parse("2006-01-02", parts[2])
	if err != nil {
		log.Fatalf("Failed to parse date for file %s. Check the format.\n", filename)
	}

	return parts[0], parts[1], date
}
