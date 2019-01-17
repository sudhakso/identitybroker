package api

import (
	"log"
	"time"
	"errors"
	
	"golang.org/x/net/context"
	prov "github.com/identitybroker/pkg/provider"
	mapper "github.com/identitybroker/api/mapper"
	rpcgen "github.com/identitybroker/api/_generated"
)

// A trace utility function
func Trace(msg string) func() {
	// print regular trace
	start := time.Now()
	log.Printf("Handler ", msg)
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
		return nil, err
	}

	stat, err := reg.RegisterProvider() // delegate
	if err != nil {
		log.Printf("Failed registering the Provider")
		return nil, err
	}
	return newRegistrationStatus(*stat), nil
}

func (r ResourceProviderService) DeRegisterProvider(ctx context.Context, opts *rpcgen.ProviderRegistrationOpts) (*rpcgen.RegistrationStatus, error) {
	defer Trace("DeRegisterProvider()")()
	
	rOpts := newRegisterOpts(*opts) //cast
	reg, err := prov.NewProviderRegistrar(rOpts) //factory
	if err != nil {
		log.Printf("Failed creating a Registrar : %v", err)
		return nil, err
	}
	
	stat, err := reg.DeRegisterProvider() // delegate
	if err != nil {
		log.Printf("Failed (de)registering the Provider")
		return nil, err
	}
	return newRegistrationStatus(*stat), nil
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
	return &prov.GenericProviderOpts{Dummy:"Dummy"}
}

func newProviderConnection(opts rpcgen.ProviderOpts) *prov.ProviderConnection {
	return &prov.ProviderConnection{Dummy:"Dummy"}
}

func newRegisterOpts(opts rpcgen.ProviderRegistrationOpts) *prov.RegisterOpts {
	return &prov.RegisterOpts{Dummy:"Dummy"}
}

func newStatus(opts rpcgen.RegistrationStatus) *prov.Status {
	return &prov.Status{Dummy:"Dummy"}
}

func newRegistrationStatus(opts prov.Status) *rpcgen.RegistrationStatus {
	return &rpcgen.RegistrationStatus{Error: true, OriginalError: "", ProviderId: "", ProviderNamespace: ""}
}
