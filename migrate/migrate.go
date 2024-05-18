package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dcConn := db.NewDB()
	defer fmt.Println("Successfully migrated")
	defer db.CloseDB(dcConn)
	dcConn.AutoMigrate(&model.User{}, &model.Task{})

}
