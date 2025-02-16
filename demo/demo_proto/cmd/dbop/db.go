package main

import (
	"github.com/joho/godotenv"
	"github.com/persue1/CloudWego/demo/demo_proto/biz/dal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	//CURD

	// Create
	// mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "jfiojffjsoij"})

	// Update
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "22222222")

	// Read
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)

	// fmt.Printf("row: %+v\n", row)

	// Delete
	// mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})

	// mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})
}
