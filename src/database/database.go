package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver
	//ORMs and tools to handle database
	//"https://github.com/volatiletech/sqlboiler/"
	//"https://github.com/prisma/prisma-client-go"
	//"gorm.io/gorm"
	//"xorm.io/xorm"
)

// Connect to database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
