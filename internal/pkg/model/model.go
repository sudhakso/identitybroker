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
	TypeName 	string			`gorm:"primary_key"`
	Path	 	string
	Version 	string
	ProviderID  int
}

type Credential struct {
	gorm.Model
	ProviderID	uint
	AccessKey	string
}

type Provider struct {
	gorm.Model
	
	Name 			string			`gorm:"primary_key"`
	Type 			string			`gorm:"type:text"`
	State			string
	Url				string
	
	Credential  	Credential		
	// one-to-many
	ResourceTypes 	[]ResourceType
}

var AutoMigrateTables = []interface{}{
	&Provider{},
	&Credential{},
	&ResourceType{},
}
