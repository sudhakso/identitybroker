package backend

import (
		"log"
		"sync"
		
		"github.com/jinzhu/gorm"
		"github.com/identitybroker/configs"
		"github.com/identitybroker/internal/pkg/model"
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
//type Orm struct {
//	gorm.DB
//}

type DBConfigOpts struct {
	Username string
	Password string
	Hostname string
	Port     string
	DBName   string
	BackendType DBBackendType
}

type Orm struct {
	DB *gorm.DB	
}

type OrmConfig struct {
	config DBConfigOpts
}

func init() {
	// Build supported drivers map
	Drivers = make(map[string]func (DBConfigOpts) (*gorm.DB, error))
}

func (o *Orm) Close() {
	defer o.DB.Close()
}

func (o *OrmConfig) OpenDBFromConfig() (*Orm, error) {
	db, err := Drivers[string(o.config.BackendType)](o.config)		
	if err != nil {
		log.Fatal("Failed to open db connection %v", err)
	}
	return &Orm{DB: db}, err
}

func NewOrmConfig(opt DBConfigOpts) (*OrmConfig) {
	return 	&OrmConfig{config: opt}
}

func InitializeDBFromRuntimeConfig(models interface{}) (error) {
	dbConfig := DBConfigOpts{Username: configs.BootConfig.DB_User,
		Password: 		configs.BootConfig.DB_Password,
		Hostname: 		configs.BootConfig.DB_Host,
		Port: 			configs.BootConfig.DB_Port,
		DBName:			configs.BootConfig.DB_Name,
		BackendType: 	DBBackendType(configs.BootConfig.DB_Type),
	}
    // Do once
	singletonOnce.Do(func(){
		log.Printf("Loading the driver %s", string(dbConfig.BackendType)) 	
		db, err := Drivers[string(dbConfig.BackendType)](dbConfig)		
		if err != nil {
			log.Fatal("Failed to connect to the db %v", err)
		}
		log.Printf("Loading tables %v", models)
		// clean start always, Only for testing
		db.Debug().DropTableIfExists(&model.Provider{})
		db.Debug().DropTableIfExists(&model.Credential{})
		db.Debug().DropTableIfExists(&model.ResourceType{})
		//create table
		//db.Debug().CreateTable(models)
		// Create
		db.Debug().CreateTable(&model.Credential{})
		db.Debug().CreateTable(&model.ResourceType{})
		db.Debug().CreateTable(&model.Provider{})
		// Migrate	
		db.Debug().AutoMigrate(&model.Credential{})
		db.Debug().AutoMigrate(&model.ResourceType{})
		db.Debug().AutoMigrate(&model.Provider{})})
	return nil
}