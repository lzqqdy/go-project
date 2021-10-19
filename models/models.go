package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/spf13/viper"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

// InitMysql Setup initializes the database instance
func InitMysql() {
	var err error
	//配置MySQL连接参数
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	Dbname := viper.GetString("mysql.dbname")
	charset := viper.GetString("mysql.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, Dbname, charset)

	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
