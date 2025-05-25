package main

import (
	"fmt"
	"go-echo/db"
	"go-echo/model"
	"log"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)

	if err := dbConn.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("Successfully Migrated")
}
