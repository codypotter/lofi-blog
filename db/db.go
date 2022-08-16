package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

var (
	ErrConnection = errors.New("db connection error")
	ErrNotFound   = errors.New("db error not found")
)

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title" gorm:"unique"`
	Markup    string    `json:"markup"`
	Category  string    `json:"category"`
	Upvotes   uint      `json:"upvotes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func Connect() {
	dbPort, err := strconv.Atoi(os.Getenv("RDS_PORT"))
	if err != nil {
		log.Panic("failed to read port env")
	}
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("RDS_USERNAME"),
		os.Getenv("RDS_PASSWORD"),
		os.Getenv("RDS_HOSTNAME"),
		dbPort,
		os.Getenv("RDS_DB_NAME"),
	)
	conn, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))
	if err != nil {
		log.Panic("failed to connect database")
	}
	err = conn.AutoMigrate(&Post{})
	if err != nil {
		log.Panicf("failed to automigrate database: %v", err)
	}

	readPosts()
}

func readPosts() {
	outputDirRead, _ := os.Open("./posts")

	postFiles, err := outputDirRead.ReadDir(0)
	if err != nil {
		log.Fatal("failed to read posts directory")
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
			log.Fatalf("error reading file %s\n", postFileName)
		}

		// parse markdown file into html
		parser := parser.New()
		html := markdown.ToHTML(md, parser, nil)

		// parse filename into post metadata
		title, category, date := parseFileName(postFileName)

		AddPost(title, string(html), category, date)
	}
}

func parseFileName(filename string) (string, string, time.Time) {
	parts := strings.Split(filename, "$")

	date, err := time.Parse("2006-01-02", parts[2])
	if err != nil {
		log.Fatalf("failed to parse date for file %s. check the format.\n", filename)
	}

	return parts[0], parts[1], date
}
