package db_operation

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const conn = "root:smoty@tcp(172.20.0.3:3306)/score?charset=utf8&parseTime=True&loc=Local"

type Chat struct {
	gorm.Model
	Text string
}

func db_connect() error {
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("dbInit_users失敗: %w", err)
	}
	db.AutoMigrate(&Chat{})
	return nil
}

func db_addition(Text string) {
	db_connect()
}
