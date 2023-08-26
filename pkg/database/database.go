package database

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func DatabaseInit(c *gin.Context) {
	db, err := sql.Open("sqlite3", "database/image.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建图像表（如果不存在）
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS images (
			id INTEGER PRIMARY KEY,
			original_name TEXT NOT NULL,
			modified_name TEXT NOT NULL,
			image_path TEXT NOT NULL,
			unique_identifier TEXT NOT NULL UNIQUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}

}
