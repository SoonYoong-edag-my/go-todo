package controllers

import (
	"github.com/go-rest/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	dsn := "host=localhost user=zacktest password=zacktest dbname=gotodo port=5434 sslmode=disable TimeZone=Asia/Kuala_Lumpur"
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connection to DB success !!")

	// Migrate the schema
	if err = dbConnect.AutoMigrate(&model.Todo{}); err != nil {
		panic("failed to create schema")
	}

	log.Println("Auto Migrate DB success !!")
	db = dbConnect
}
