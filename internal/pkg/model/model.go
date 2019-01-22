package model

import (
	"github.com/jinzhu/gorm"
)

const (
	TYPE_USER string 	= "Users"
	TYPE_GROUP string 	= "Groups"
	TYPE_APPLICATION string = "Applications" 
)

type ResourceType struct {
	TypeName 	string
	Path	 	string
	Version 	string
}

type Credential struct {
	gorm.Model
	
	AccessKey	string	
}

type Provider struct {
	gorm.Model
	
	Name 			string		`gorm:"PRIMARY_KEY;UNIQUE"`
	Type 			string		`gorm:"type:text"`
	State			string
	Url				string
	Credential  	Credential
	ResourceTypes 	[]ResourceType
}

var AutoMigrateTables = []interface{}{
	&Provider{},
	&Credential{},
	&ResourceType{},
}
