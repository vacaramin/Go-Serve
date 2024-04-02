package initializers

import (
	"url-shortner/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DB_init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("url_shortener.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.Url{})

	// Create a sample URL
	DB.Create(&models.Url{URL: "https://example.com", ShortenedURL: "abc123"})

	// Read the URL
	var url models.Url
	DB.First(&url, 1)                             // find URL with integer primary key
	DB.First(&url, "shortened_url = ?", "abc123") // find URL with shortened URL "abc123"

	// Update the URL
	DB.Model(&url).Updates(models.Url{URL: "https://example.com/update", ShortenedURL: "xyz789"})

	// Delete the URL
	DB.Delete(&url, 1)
}
