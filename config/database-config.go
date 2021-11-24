package config

import (
	"github.com/sholehbaktiabadi/go-api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Create Connection
func ConnectDatabase() *gorm.DB {
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")

	dsn := "host=localhost user=postgres password=root dbname=db_go_proper port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect databases")
	}
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

//Close connection
func CloneConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close databse connection")
	}
	dbSql.Close()
}
