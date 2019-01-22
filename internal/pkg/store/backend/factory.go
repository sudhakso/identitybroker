package backend

import (
		"log"
		"sync"
		
		"github.com/jinzhu/gorm"
)

var (
		singletonDB   *gorm.DB
		singletonOnce sync.Once
)


// Regsitered drivers
var Drivers map[string]func (DBConfigOpts) (*gorm.DB, error)

type DBBackendType string

const (
	MYSQL 		DBBackendType = "mysql"
	SQLITE3 	DBBackendType = "sqlite3"
	POSTGRES	DBBackendType = "postgres"
)

// backend DB connection : gorm db rds
type Orm struct {
	DB *gorm.DB
}

type DBConfigOpts struct {
	Username string
	Password string
	Hostname string
	Port     string
	DBName   string
	BackendType DBBackendType
}

func init() {
	Drivers = make(map[string]func (DBConfigOpts) (*gorm.DB, error))	
}

func NewDBFromConfig(opt DBConfigOpts, models interface{}) (*Orm, error) {
	
	singletonOnce.Do(func(){
		log.Printf("Loading the driver %s", string(opt.BackendType)) 	
		db, err := Drivers[string(opt.BackendType)](opt)		
		if err != nil {
			log.Fatal("Failed to connect to the db %v", err)
		}
		log.Printf("Loading tables %v", models)
		//load tables once
		db = db.AutoMigrate(models)
		singletonDB = db})
	return &Orm{DB: singletonDB.Debug()}, nil
}