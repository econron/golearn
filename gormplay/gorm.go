package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}  
  
func main() {
// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "docker:docker@tcp(gosql:3306)/test_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	if err != nil {
		// fmt.Println(err)
	}

	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D42", Price: 100})

}