package provider

import (
	"log"
	
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

type Status struct { //RegistrationStatus
	Dummy string
	RegistrationId 		string
	RegistrationName  	string
	CurrentState		string	
	NewlyCreated		bool
	RequestName			string
	Errored				bool
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
	// db config
	dbConfig := backend.DBConfigOpts{Username: configs.BootConfig.DB_User,
		Password: 		configs.BootConfig.DB_Password,
		Hostname: 		configs.BootConfig.DB_Host,
		Port: 			configs.BootConfig.DB_Port,
		DBName:			configs.BootConfig.DB_Name,
		BackendType: 	backend.DBBackendType(configs.BootConfig.DB_Type),
	}
	// TBD error
	// Always gets a singleton instance of DB and wraps it with ORM interface
	db, err := backend.NewDBFromConfig(dbConfig, model.AutoMigrateTables)
	if err != nil {
		return nil, err
	}

	// Success
	return &ProviderRegistrationFactory{
		Options: opts,
		orm: store.ProviderORM{
			Provider: model.Provider{
				Name: opts.Namespace,
				Type: opts.ProviderType,
				State: "Active",
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
	// Connect to the provider
	// Store connection parameters
	if _, err := r.orm.Store(); err != nil {
		stat := &Status{RequestName: "RegisterProvider", Errored: true}
		return stat, err
	}
	
	//Stored, get hidden values
	prov := model.Provider{Name: r.Options.Namespace, Type: r.Options.ProviderType}
	result, err := r.orm.Get(prov)
	if err != nil {
		//erro in getting
		stat := &Status{RequestName: "RegisterProvider", Errored: true}
		return stat, err
	}
	
	v, ok := result.(model.Provider);
	if !ok {
		//erro in getting
		stat := &Status{RequestName: "RegisterProvider", Errored: true}
		return stat, err
	} 
	stat := &Status{RequestName: "RegisterProvider",
		 RegistrationId		: 	string(v.ID),
		 RegistrationName	: 	v.Name,
		 CurrentState		: 	v.State,
	}
	return stat, nil
}

func (r *ProviderRegistrationFactory) DeRegisterProvider() (*Status, error) {
	return nil, nil
}