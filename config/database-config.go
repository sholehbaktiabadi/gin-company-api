package config

import (
	"fmt"
	"log"

	"github.com/sholehbaktiabadi/go-api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Create Connection
func ConnectDatabase() *gorm.DB {
	env, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read env vars:", err)
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect databases")
	}
	db.AutoMigrate(&entity.Company{}, &entity.User{})
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
