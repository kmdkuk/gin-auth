package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kmdkuk/gin-auth/model"
)

var (
	db *gorm.DB
)

const (
	// データベース
	dialect = "mysql"

	// TODO: ユーザ名とパスワードは.envなど環境変数から読み込む。
	// ユーザー名
	dbUser = "root"

	// プロトコル
	dbProtocol = "tcp(db:3306)"

	// DB名
	dbName = "go_auth"
)

// Init is initialize db from main function
func Init() error {
	dbPass := os.Getenv("MYSQL_ROOT_PASSWORD")
	connectTemplate := "%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local"
	connect := fmt.Sprintf(connectTemplate, dbUser, dbPass, dbProtocol, dbName)
	var err error
	db, err = gorm.Open(dialect, connect)
	if err != nil {
		return err
	}
	if err := autoMigration(); err != nil {
		return err
	}
	return nil
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() error {
	db.AutoMigrate(&model.User{})
	return db.Error
}
