package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title   string
	Content string
}

func main() {
	dsn := "example_user:example_password@tcp(127.0.0.1:3306)/example_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Миграция (создание таблицы)
	db.AutoMigrate(&Note{})

	// Создание новой заметки
	db.Create(&Note{Title: "Need to do something today", Content: "First thing to do"})
}
