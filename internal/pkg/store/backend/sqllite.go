package backend

import (
	"log"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	Drivers[string(SQLITE3)] = OpenMySqlite3DB 	
}


func OpenMySqlite3DB(opts DBConfigOpts) (*gorm.DB, error) {
	connectionStr := "/tmp/" + opts.DBName + ".db"
	log.Printf("Connecting to database %s", connectionStr)
	
	return gorm.Open(string(opts.BackendType), connectionStr)
}