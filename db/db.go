package db

import (
	"fmt"
	"os"

	"github.com/cla-github/file_manager/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}
}

var Users = []types.User{
	{
		Id:       "kfjie939",
		Username: "cla",
		Password: "$2a$14$ZQsaN4aXs8bYUfaH4pgXmeCupUdXGwzsiVkvuqQjf8xhQFZdA3YrS",
		Email:    "cla@outlook.com",
	},
}
