package provider

import (
	"log"
	"errors"

	"github.com/identitybroker/configs"
	"github.com/identitybroker/internal/pkg/model"
	"github.com/identitybroker/internal/pkg/store"
	"github.com/identitybroker/internal/pkg/store/backend"
)

// moves to ORM layer
type ResourceReaderOpts struct { // GetResourceOpts //placeholder
  Dummy string  
  ResourceUrl string
  ResourceQuery map[string]string
}

// moves to ORM layer
type GenericProviderOpts struct { //ProviderOpts //placeholder
  Dummy 			string
  Namespace			string
  ProviderId		string
}

// moves to ORM layer
type ProviderResource struct {
	Id    string
	Name  string
	Kind  string
	State string
	// test
	Dummy string
}

// moves to ORM layer
type ProviderResourceType struct {
	Name  string
	Location string
	// test
	Dummy string
}

// moves to ORM layer
type ProviderResourceTypes struct {
    Types []*ProviderResourceType
}

type ProviderResources struct {
	Resources []*ProviderResource	
}

// ORM layer
type ProviderORM struct {
	model.Provider
	BaseOrm *backend.Orm
}

// gets a new orm
func NewProviderORM(filter *ProviderResource) (*ProviderORM, error) {
		// db config
	dbConfig := backend.DBConfigOpts{Username: configs.BootConfig.DB_User,
		Password: 		configs.BootConfig.DB_Password,
		Hostname: 		configs.BootConfig.DB_Host,
		Port: 			configs.BootConfig.DB_Port,
		DBName:			configs.BootConfig.DB_Name,
		BackendType: 	backend.DBBackendType(configs.BootConfig.DB_Type),
	}
	// Always gets a singleton instance of DB and wraps it with ORM interface
	db, err := backend.NewOrmConfig(dbConfig).OpenDBFromConfig()
	if err != nil {
		return nil, err
	}
	// Returns unitialized ORM with DB reference
	return &ProviderORM{
		Provider: model.Provider{
					Name	: filter.Name,
					Type	: filter.Kind,
					State	: filter.State,
		},
		BaseOrm: db}, nil
}

func (p *ProviderORM) Store() (alreadyExists bool, err error) {
	// Stores only Provider
	log.Printf("Inserting provider record %s", p.Provider.Name)
	// Insert a new record
	p.BaseOrm.DB.Debug().Create(&p.Provider)
	log.Printf("Done inserting provider record %s", p.Provider.Name)
	return
}

func (p *ProviderORM) Get() (store.Value, error) {
	var provRecord model.Provider
	
	q := &model.Provider{Name: p.Name}
	p.BaseOrm.DB.Debug().Where(q).Find(&provRecord)
	log.Println(provRecord)
	return provRecord, nil
}

func (p *ProviderORM) Close() {
	// Do anything?
	p.BaseOrm.Close()
}

func (p *ProviderORM) Update(option store.Value) (error) {
	v, ok := option.(UpdateOpts)
	if !ok {
		log.Println("Failed updating record")
		return errors.New("Failed updating records")
	}

	log.Println(v)
	
	q := &model.Provider{Name: p.Name}
	var prov model.Provider
	p.BaseOrm.DB.Debug().Where(q).First(&prov)

	// update credentials if specified
	if len(v.ApiKey) > 0 {
		var cred model.Credential
		p.BaseOrm.DB.Debug().Model(&prov).Related(&cred)
		log.Println(cred)
		p.BaseOrm.DB.Debug().Model(&cred).Update("AccessKey", v.ApiKey)
		log.Println(cred)
	}
	
	// Update domain URl if specified
	// update provider record the last
	if len(v.DomainUrl) > 0 {
		p.BaseOrm.DB.Debug().Model(&prov).Update("Url", v.DomainUrl)		
	}
	return nil
}

// Reader
type ResourceReader interface {
	GetResource(*ResourceReaderOpts) (*ProviderResource, error)
	ListResource(*ResourceReaderOpts) (*ProviderResources, error)
} 

type ReadOnlyResourceProvider struct {
	// provider opts
	providerConnOpts 		*model.Provider
	// query options
	AllowedAccessType 		string
	ResourceUrl		  		string
	ResourceQueryFilter     map[string]string
}

func NewResourceReaderFromOptions(opt *GenericProviderOpts) (ResourceReader, error) {
	log.Printf("Creating resource reader for %s : %s", opt.ProviderId, opt.Namespace)
   
    // Build the ORM
    pOrm, err := NewProviderORM(&ProviderResource{Name: opt.Namespace, Id: opt.ProviderId})
	if err != nil {
		return nil, err
	}
	// get the provider resource
	val, err := pOrm.Get()
	if err != nil {
		return nil, err
	}
	// Cast to Provider
	v, ok := val.(model.Provider)
	if !ok {
		return nil, errors.New("DB error")
	}
	// Load the connection object
	rr := ReadOnlyResourceProvider{providerConnOpts : &v}
	return &rr, nil
}
	
func (p *ReadOnlyResourceProvider) GetResource(opts *ResourceReaderOpts) (*ProviderResource, error) {
	return nil, nil
} 

func (p *ReadOnlyResourceProvider) ListResource(opts *ResourceReaderOpts) (*ProviderResources, error) {
	return nil, nil
}


// Type reader
type ResourceTypeReader interface {
	GetResourceTypes() (*ProviderResourceTypes, error)
}

type ReadOnlyResourceTypeProvider struct {
	Options *GenericProviderOpts
}

func NewResourceTypeReaderFromOptions(opt *GenericProviderOpts) (ResourceTypeReader, error) {
	return nil, nil
}

func (p *ReadOnlyResourceTypeProvider) GetResourceTypes() (*ProviderResourceTypes, error) {
	return nil, nil
} 