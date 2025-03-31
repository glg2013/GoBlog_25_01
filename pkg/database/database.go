// Package database 数据库相关
package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"goblog/pkg/logger"
	"time"
)

// DB 数据库对象
var DB *sql.DB

// Initialize 初始化数据库
func Initialize() {
	initDB()
	createTables()
}

func initDB() {

	var err error

	config := mysql.Config{
		User:                 "homestead",
		Passwd:               "secret",
		Addr:                 "192.168.56.56:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	// 打印信息
	// fmt.Println("Connecting to database ...", config.FormatDSN())
	// Connecting to database ... homestead:secret@tcp(192.168.56.56:3306)/goblog?checkConnLiveness=false&maxAllowedPacket=0

	// 准备数据库连接池
	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	DB.SetMaxIdleConns(25)
	// 设置每个连接的过期时间
	DB.SetConnMaxLifetime(time.Minute * 5)

	// 尝试连接，失败会报错
	err = DB.Ping()
	logger.LogError(err)
}

// 创建博客相关的表
func createTables() {
	createArticlesSQL := `
    	CREATE TABLE IF NOT EXISTS articles(
    	    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    	    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    	    body longtext COLLATE utf8mb4_unicode_ci
    	);
    `

	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
