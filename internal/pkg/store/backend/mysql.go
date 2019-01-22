package backend

import (
	"log"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	Drivers[string(MYSQL)] = OpenMySqlDB 	
}


func OpenMySqlDB(opts DBConfigOpts) (*gorm.DB, error) {
	connectionStr := opts.Username + ":" + opts.Password + "@tcp(" + opts.Hostname + ":" + opts.Port + ")/" + opts.DBName + "?parseTime=true"
	log.Printf("Connecting to database %s", connectionStr)
	
	return gorm.Open(string(opts.BackendType), connectionStr)
}