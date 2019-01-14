package api

import (
	"log"
	"time"
	
	"golang.org/x/net/context"
	"github.com/identitybroker/pkg/provider"
)

// A trace utility function
func Trace(msg string) func() {
	// print regular trace
	start := time.Now()
	log.Printf("Handler ", msg)
	return func() {log.Printf("Finished handler %s %s\n", msg, time.Since(start))}
}

// gRPC Server handler
type ResourceProviderService struct {}

func (r *ResourceProviderService) RegisterProvider(ctx context.Context, opts *ProviderRegistrationOpts) (*RegistrationStatus, error) {
	defer Trace("RegisterProvider()")()
	
	rOpts := newRegisterOpts(*opts) //cast
	reg, err := NewProviderRegistrar(&rOpts) //factory
	if err != nil {
		log.Printf("Failed creating a Registrar : %v", err)
		return nil, err
	}

	stat, err := reg.RegisterProvider() // delegate
	if err != nil {
		log.Printf("Failed registering the Provider")
		return nil, err
	}
	return &newRegistrationStatus(stat), nil
}

func (r *ResourceProviderService) DeRegisterProvider(ctx context.Context, opts *ProviderRegistrationOpts) (*RegistrationStatus, error) {
	defer Trace("DeRegisterProvider()")()
	
	rOpts := newRegisterOpts(*opts) //cast
	reg, err := NewProviderRegistrar(&rOpts) //factory
	if err != nil {
		log.Printf("Failed creating a Registrar : %v", err)
		return nil, err
	}
	
	stat, err := reg.DeRegisterProvider() // delegate
	if err != nil {
		log.Printf("Failed (de)registering the Provider")
		return nil, err
	}
	return &newRegistrationStatus(stat), nil
}

func (r *ResourceProviderService) GetResource(ctx context.Context, opts *GetResourceOpts) (*Resource, error) {
	defer Trace("GetResource()")()
	
	if opts == nil {
		log.Printf("Unknown options")
		return nil, errors.Errorf("Resource option cannot be nil")
	}
	
	if p := opts.Provider; p != nil {
		reader, err := NewResourceReaderFromOptions(newGenericProviderOpts(p))
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
		return res, err
	}
}

func (r *ResourceProviderService) ListResource(ctx context.Context, opts *GetResourceOpts) (*Resources, error) {
	defer Trace("ListResource()")()
	
	if opts == nil {
		log.Printf("Unknown options")
		return nil, errors.Errorf("Resource option cannot be nil")
	}
	
	if p := opts.Provider; p != nil {
		reader, err := NewResourceReaderFromOptions(newGenericProviderOpts(p))
		if err != nil {
			log.Printf("Could not parse provider : %v", err)
			return nil, err
		}
		
		//call into provider returned
		res, err := reader.ListResource(ResourceReaderOpts(*opts))
		if err != nil {
			log.Printf("Could list resource from the provider : %v", err)
			return nil, err
		}
		return res, err
	}	
}

func (r *ResourceProviderService) GetResourceTypes(ctx context.Context, opts *ProviderOpts) (*ResourceTypes, error) {
	defer Trace("GetResourceTypes()")()
		
	if opts == nil {
		log.Printf("Unknown options")
		return nil, errors.Errorf("provider option cannot be nil")
	}
	if opts != nil {
		rtype, err := NewResourceTypeReaderFromOptions(opts)
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
		return types, err
	}	
}

func (r *ResourceProviderService) AssignResourceLink(context.Context, *ResourceBindingOpts) (*ResourceBindings, error) {
	defer Trace("AssignResourceLink()")()
	return nil, nil
}

func (r *ResourceProviderService) RemoveResourceLink(context.Context, *ResourceBindingOpts) (*ResourceBindings, error) {
	defer Trace("RemoveResourceLink()")()
	return nil, nil	
}

// TBD: Move to Serializer
func newResourceReaderOpts(opts GetResourceOpts) *ResourceReaderOpts {
	return &ResourceReaderOpts{dummy:"dummy"}
}

func newGenericProviderOpts(opts ProviderOpts) *GenericProviderOpts {
	return &GenericProviderOpts{dummy:"dummy"}
}

func newProviderConnection(opts ProviderOpts) *ProviderConnection {
	return &ProviderConnection{dummy:"dummy"}
}

func newRegisterOpts(opts ProviderRegistrationOpts) *RegisterOpts {
	return &RegisterOpts{dummy:"dummy"}
}

type newStatus(opts RegistrationStatus) *Status {
	return &Status{dummy:"dummy"}
}
