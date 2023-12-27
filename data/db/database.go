package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)


// Connect connects to a database using the provided DSN.
//
// Parameters:
// - dsn: The Data Source Name (DSN) used to connect to the database.
//
// Returns:
// - *gorm.DB: A pointer to the connected gorm.DB object.
// - error: An error object indicating any connection errors that occurred.
func Connect(dsn string) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		dsn = fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", dsn)
		fmt.Println(dsn)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})
	return db, err
}
