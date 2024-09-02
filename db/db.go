package db

import (
	"blog-api/models"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){
	 dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=verify-full sslrootcert=ca.pem",
    //     dbHost, dbPort, dbUser, dbPassword, dbName)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require sslrootcert=db/ca.pem",
    dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	// DB, err = sql.Open("postgres", connStr)
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Blog{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	//    if err := DB.Ping(); err != nil {
    //     log.Fatalf("Failed to ping database: %v", err)
    // }

    log.Println("Database connection successful")
}