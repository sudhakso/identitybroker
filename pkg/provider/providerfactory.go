package provider

import (
	"log"
	"fmt"
	"errors"
	
	"github.com/identitybroker/configs"
	"github.com/identitybroker/internal/pkg/store"
	"github.com/identitybroker/internal/pkg/store/backend"
	"github.com/identitybroker/internal/pkg/model"
)

type ProviderConnection struct { //ProviderOpts
  Dummy string
}

type RegisterOpts struct { //ProviderRegistrationOpts
	Dummy 			string
	Namespace		string
	ProviderType 	string
	DomainUrl		string
	ApiKey			string
}

type UpdateOpts struct { //ProviderRegistrationOpts
	DomainUrl		string
	ApiKey			string
}

type Status struct { //RegistrationStatus
	Dummy string
	// Request info
	RequestName			string
	// Provider data
	ProviderId			string
	ProviderName		string
	Locations			[]string
	// Status
	State				string
	StateRef			uint
	// Errors
	ErrorCode			int
	Description			string
}

type Registrar interface {
	RegisterProvider() (*Status, error) 
	DeRegisterProvider() (*Status, error)
} 

// TBD
type ProviderRegistrationFactory struct {
	Options *RegisterOpts
	orm 	store.ModelReaderWriter
}

func NewProviderRegistrar(opts *RegisterOpts) (Registrar, error) {
	log.Printf("Registration factory for %s", opts.Namespace)
	
	//TBD: move to start up
	backend.InitializeDBFromRuntimeConfig(model.AutoMigrateTables)
	
	// db config
	dbConfig := backend.DBConfigOpts{Username: configs.BootConfig.DB_User,
		Password: 		configs.BootConfig.DB_Password,
		Hostname: 		configs.BootConfig.DB_Host,
		Port: 			configs.BootConfig.DB_Port,
		DBName:			configs.BootConfig.DB_Name,
		BackendType: 	backend.DBBackendType(configs.BootConfig.DB_Type),
	}
	// TBD error
	db, err := backend.NewOrmConfig(dbConfig).OpenDBFromConfig()
	if err != nil {
		return nil, err
	}

	// Success
	return &ProviderRegistrationFactory{
		Options: opts,
		orm: &ProviderORM {
			Provider: model.Provider{
				Name: opts.Namespace,
				Type: opts.ProviderType,
				State: "Pending",
				Url: opts.DomainUrl,
				Credential: model.Credential{AccessKey: opts.ApiKey},
				// TBD : Move to Provider specific driver
				ResourceTypes: []model.ResourceType{
					model.ResourceType{TypeName: model.TYPE_USER, Path: "/resources/"+model.TYPE_USER},
					model.ResourceType{TypeName: model.TYPE_GROUP, Path: "/resources/"+model.TYPE_GROUP},
					model.ResourceType{TypeName: model.TYPE_APPLICATION, Path: "/resources/"+model.TYPE_APPLICATION},
				},
			}, 
			BaseOrm: db},
	}, nil
}

func (r *ProviderRegistrationFactory) RegisterProvider() (*Status, error) {
	log.Printf("Registering the Provider %s", r.Options.Namespace)
	// Store connection parameters
	if _, err := r.orm.Store(); err != nil {
		stat := &Status{
			RequestName	: "RegisterProvider",
			ProviderId	: "",
			ProviderName: r.Options.Namespace,
			State		: "Error",
			ErrorCode	: 102,
			Description	: fmt.Sprintf("Failed to store provider details in database, Error = %v", err),
		} 
		log.Printf("Error : %v", stat)
		return stat, err
	}
	
	// Get provider details	
	result, err := r.orm.Get()
	if err != nil {
		stat := &Status{
			RequestName	: "RegisterProvider",
			ProviderId	: "",
			ProviderName: r.Options.Namespace,
			State		: "Error",
			ErrorCode	: 103,
			Description	: fmt.Sprintf("Failed to retrieve provider details from database, Error = %v", err),
		} 
		log.Printf("Error : %v", stat)
		return stat, err
	}
	// Map Results
	v, ok := result.(model.Provider);
	if !ok {
			stat := &Status {
			RequestName	: "RegisterProvider",
			ProviderId	: "",
			ProviderName: r.Options.Namespace,
			State		: "Error",
			ErrorCode	: 104,
			Description	: fmt.Sprintf("Error mapping data from the database for %s", r.Options.Namespace),
			}
		log.Printf("Error : %v", stat)
		return stat, errors.New(fmt.Sprintf("Error mapping data from the database for %s", r.Options.Namespace))
	}
	
	// TBD: Spawn asynchronous task to update Provider
	
	// Return valid results to the caller
	stat := &Status {
		RequestName		: "RegisterProvider",
		ProviderId		: string(v.ID),
		ProviderName	: v.Name,
		State			: v.State,
		ErrorCode		: 0,
		Description		: fmt.Sprintf("Succesfully registered the provider %s", r.Options.Namespace),
	}	
	return stat, nil
}

func (r *ProviderRegistrationFactory) DeRegisterProvider() (*Status, error) {
	return nil, nil
}
