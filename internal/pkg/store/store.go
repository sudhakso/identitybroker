package store

import (
	"log"
	"errors"
	
	"github.com/identitybroker/internal/pkg/model"
	"github.com/identitybroker/internal/pkg/store/backend"
)

type Value interface {}

type ModelReaderWriter interface {
	
	Store() (bool, error)
	Get(query Value) (Value, error)
	Exec(query string) (Value, error)
}

type ProviderORM struct {
	model.Provider
	BaseOrm *backend.Orm
}

func (p ProviderORM) Store() (alreadyExists bool, err error) {
	// Stores only Provider
	
	//Check if the provider already exists
	log.Printf("Checking if Provider record with Name : %s exists", p.Provider.Name)
	if !p.BaseOrm.DB.Debug().NewRecord(&p.Provider) {
		alreadyExists = true
		err = nil
		return
	}
	log.Printf("Not Found! Attempting to create a provider record %v", p.Provider)
	// Create and test the new record
	p.BaseOrm.DB.Debug().Create(&p.Provider)
	
	//Check if it was created
	var prov model.Provider
	if p.BaseOrm.DB.Debug().Where(&model.Provider{Name: p.Provider.Name}).First(&prov).RecordNotFound() {
		err = errors.New("Record not created")
		log.Printf("Provider record creation Failed %s", p.Provider.Name)
		return
	} else {
		log.Printf("Provider record created %v", prov)
	}
	return
}

func (p ProviderORM) Get(val Value) (Value, error) {
	var provRecord model.Provider
	// Gets only Provider
	v, ok := val.(model.Provider)
	if ok {
		p.BaseOrm.DB.Where(&v).First(&provRecord)
		return provRecord, nil
	} else {
		return nil, errors.New("Type error")
	}
}

func (p ProviderORM) Exec(query string) (Value, error) {
	return nil, nil
}
