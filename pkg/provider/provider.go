package provider

import (
)

type ResourceReaderOpts struct { // GetResourceOpts //placeholder
  dummy string  
}
type GenericProviderOpts struct { //ProviderOpts //placeholder
  dummy string
}

// Reader
type ResourceReader interface {
	GetResource(*ResourceReaderOpts) (*Resource, error)
	ListResource(*ResourceReaderOpts) (*Resources, error)
} 

type ReadOnlyResourceProvider struct {
	Options *GenericProviderOpts
}

func NewResourceReaderFromOptions(opt *GenericProviderOpts) (*ResourceReader, error) {
	return nil, nil
}
	
func (p *ReadOnlyResourceProvider) GetResource(opts *ResourceReaderOpts) (*Resource, error) {
	return nil, nil
} 

func (p *ReadOnlyResourceProvider) ListResource(opts *ResourceReaderOpts) (*Resources, error) {
	return nil, nil
}


// Type reader
type ResourceTypeReader interface {
	GetResourceTypes() (*ResourceTypes, error)
}

type ReadOnlyResourceTypeProvider struct {
	Options *GenericProviderOpts
}

func NewResourceTypeReaderFromOptions(opt *GenericProviderOpts) (*ResourceTypeReader, error) {
	return nil, nil
}

func (p *ReadOnlyResourceTypeProvider) GetResourceTypes() (*ResourceTypes, error) {
	return nil, nil
} 