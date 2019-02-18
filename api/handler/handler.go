package api

import (
	"log"
	"time"
	"errors"
	
	"golang.org/x/net/context"
	
	prov "github.com/identitybroker/pkg/provider"
	mapper "github.com/identitybroker/api/mapper"
	rpcgen "github.com/identitybroker/api/_generated"
	model "github.com/identitybroker/internal/pkg/model"
)

// A trace utility function
func Trace(msg string) func() {
	// print regular trace
	start := time.Now()
	log.Printf("Handler %s", msg)
	return func() {log.Printf("Finished handler %s %s\n", msg, time.Since(start))}
}

// gRPC Server handler
type ResourceProviderService struct {
	ResMapper mapper.ResourceSerializer
}

func (r ResourceProviderService) RegisterProvider(ctx context.Context, opts *rpcgen.ProviderRegistrationOpts) (*rpcgen.RegistrationStatus, error) {
	defer Trace("RegisterProvider()")()

	rOpts := newRegisterOpts(*opts) //cast
	reg, err := prov.NewProviderRegistrar(rOpts) //factory
	if err != nil {
		log.Printf("Failed creating a Registrar : %v", err)
		// Error
		return newRegistrationStatusFromError(err), err
	}

	stat, err := reg.RegisterProvider() // delegate
	if err != nil {
		// Error
		return newRegistrationStatus(stat), err
	}
	// Success
	return newRegistrationStatus(stat), nil
}

func (r ResourceProviderService) DeRegisterProvider(ctx context.Context, opts *rpcgen.ProviderRegistrationOpts) (*rpcgen.RegistrationStatus, error) {
	defer Trace("DeRegisterProvider()")()
	
	rOpts := newRegisterOpts(*opts) //cast
	reg, err := prov.NewProviderRegistrar(rOpts) //factory
	if err != nil {
		log.Printf("Failed creating a Registrar : %v", err)
		return newRegistrationStatusFromError(err), err
	}
	
	stat, err := reg.DeRegisterProvider() // delegate
	if err != nil {
		log.Printf("Failed (de)registering the Provider")
		return nil, err
	}
	return newRegistrationStatus(stat), nil
}

func (r ResourceProviderService) UpdateProvider(
									ctx context.Context,
									opts *rpcgen.ProviderUpdateOpts) (*rpcgen.RegistrationStatus, error) {
	defer Trace("UpdateProvider()")()
	//create providerORM resource
	orm, err := prov.NewProviderORM(newProviderFilterFromInput(opts))
	if err != nil {
		return newRegistrationStatusFromError(err), err
	}
	// Update records and related
	updateerr := orm.Update(newProviderUpdateOptsFromInput(opts))
	if updateerr != nil {
		return newRegistrationStatusFromError(updateerr), updateerr
	}
	
	//get latest provider record
	v, geterr := orm.Get()
	if geterr != nil {
		return newRegistrationStatusFromError(geterr), geterr
	}
	// cast the provider; should never Fail
	if p, ok := v.(model.Provider); ok {
		return newRegistrationStatus(&prov.Status{
										RequestName	: "UpdateProvider",
										ProviderId	: string(p.ID),
										ProviderName: p.Name,
							}), nil
		}
	// unknown error
	e := errors.New("Unknown error")
	return newRegistrationStatusFromError(e), e
}								

func (r ResourceProviderService) GetResource(ctx context.Context, opts *rpcgen.GetResourceOpts) (*rpcgen.Resource, error) {
	defer Trace("GetResource()")()
	
	if opts == nil {
		log.Printf("Unknown options")
		return nil, errors.New("Resource option cannot be nil")
	}
	
	if p := opts.Provider; p != nil {
		reader, err := prov.NewResourceReaderFromOptions(newGenericProviderOpts(*p))
		if err != nil {
			log.Printf("Could not parse provider : %v", err)
			return nil, err
		}
		
		//call into provider returned
		res, err := reader.GetResource(newResourceReaderOpts(*opts))
		if err != nil {
			log.Printf("Could get resource from the provider : %v", err)
			return nil, err
		}
		
		result, err := r.ResMapper.Serialize(res)
		if err != nil {
			log.Printf("Could get mapp resource type : %v", err)
			return nil, err
		}
		return result, err
	}
	return nil, errors.New("provider is Null")
}

func (r ResourceProviderService) ListResource(ctx context.Context, opts *rpcgen.GetResourceOpts) (*rpcgen.Resources, error) {
	defer Trace("ListResource()")()
	
	if opts == nil {
		log.Printf("Unknown options")
		return nil, errors.New("Resource option cannot be nil")
	}
	
	if p := opts.Provider; p != nil {
		reader, err := prov.NewResourceReaderFromOptions(newGenericProviderOpts(*p))
		if err != nil {
			log.Printf("Could not parse provider : %v", err)
			return nil, err
		}
		
		//call into provider returned
		res, err := reader.ListResource(newResourceReaderOpts(*opts))
		if err != nil {
			log.Printf("Could list resource from the provider : %v", err)
			return nil, err
		}
		
		result, err := r.ResMapper.SerializeCollection(res)
		if err != nil {
			log.Printf("Could get mapp resource type : %v", err)
			return nil, err
		}
		
		return result, err
	}
	return nil, errors.New("Provider context is Nil")
}

func (r ResourceProviderService) GetResourceTypes(ctx context.Context, opts *rpcgen.ProviderOpts) (*rpcgen.ResourceTypes, error) {
	defer Trace("GetResourceTypes()")()
		
	if opts == nil {
		log.Printf("Unknown options")
		return nil, errors.New("provider option cannot be nil")
	}
	if opts != nil {
		rtype, err := prov.NewResourceTypeReaderFromOptions(newGenericProviderOpts(*opts))
		if err != nil {
			log.Printf("Could not parse provider : %v", err)
			return nil, err
		}
		
		//call into provider returned
		types, err := rtype.GetResourceTypes()
		if err != nil {
			log.Printf("Could list resource types from the provider : %v", err)
			return nil, err
		}
		
		// serialize
		result := rpcgen.ResourceTypes{Types: make([]*rpcgen.ResourceType, 10, 20)}
		for _, t := range types.Types {
			atype, err :=  r.ResMapper.SerializeType(t)
			if err != nil {
				log.Printf("Error serializing object")
				return nil, err
			}
			result.Types = append(result.Types, atype)
		}
		return &result, err
	}
	return nil, errors.New("Provider option cannot be nil")
}

func (r ResourceProviderService) AssignResourceLink(context.Context, *rpcgen.ResourceBindingOpts) (*rpcgen.ResourceBindings, error) {
	defer Trace("AssignResourceLink()")()
	return nil, nil
}

func (r ResourceProviderService) RemoveResourceLink(context.Context, *rpcgen.ResourceBindingOpts) (*rpcgen.ResourceBindings, error) {
	defer Trace("RemoveResourceLink()")()
	return nil, nil	
}

// TBD: Move to Serializer
func newResourceReaderOpts(opts rpcgen.GetResourceOpts) *prov.ResourceReaderOpts {
	return &prov.ResourceReaderOpts{Dummy:"Dummy"}
}

func newGenericProviderOpts(opts rpcgen.ProviderOpts) *prov.GenericProviderOpts {
	return &prov.GenericProviderOpts{Dummy:"Dummy", Namespace:opts.Namespace, ProviderId:opts.ProviderId}
}

func newProviderConnection(opts rpcgen.ProviderOpts) *prov.ProviderConnection {
	return &prov.ProviderConnection{Dummy:"Dummy"}
}

func newRegisterOpts(opts rpcgen.ProviderRegistrationOpts) *prov.RegisterOpts {
	return &prov.RegisterOpts{
		Dummy		:"Dummy",
		Namespace	: opts.Namespace, 
		ProviderType: opts.ProviderType,
		DomainUrl	: opts.Cred.AuthUrl,
		ApiKey		: opts.Cred.ApiKey,
	}
}

func newStatus(opts rpcgen.RegistrationStatus) *prov.Status {
	return &prov.Status{Dummy:"Dummy"}
}

func newRegistrationStatus(opts *prov.Status) *rpcgen.RegistrationStatus {
	return &rpcgen.RegistrationStatus{
		Error: false,
		ProviderId: opts.ProviderId,
		ProviderNamespace: opts.ProviderName}
}

func newRegistrationStatusFromError(err error) *rpcgen.RegistrationStatus {
	return &rpcgen.RegistrationStatus{
		Error: true, 
		OriginalError: err.Error()}
}

func newProviderFilterFromInput(opts *rpcgen.ProviderUpdateOpts) *prov.ProviderResource {
	return &prov.ProviderResource{
					Id: opts.ProviderId,
					Name: opts.ProviderName,
	}
}

func newProviderUpdateOptsFromInput(opts *rpcgen.ProviderUpdateOpts) prov.UpdateOpts {
	return prov.UpdateOpts{
					ApiKey: opts.Cred.ApiKey,
					DomainUrl: opts.Cred.AuthUrl,
	}
}

