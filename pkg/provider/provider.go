package provider

import (

)

// moves to ORM layer
type ResourceReaderOpts struct { // GetResourceOpts //placeholder
  Dummy string  
}

// moves to ORM layer
type GenericProviderOpts struct { //ProviderOpts //placeholder
  Dummy string
}

// moves to ORM layer
type ProviderResource struct {
	Dummy string
	Kind  string
}

// moves to ORM layer
type ProviderResourceType struct {
	Dummy string
	Kind  string
}

// moves to ORM layer
type ProviderResourceTypes struct {
    Types []*ProviderResourceType
}

type ProviderResources struct {
	Resources []*ProviderResource	
}

// Reader
type ResourceReader interface {
	GetResource(*ResourceReaderOpts) (*ProviderResource, error)
	ListResource(*ResourceReaderOpts) (*ProviderResources, error)
} 

type ReadOnlyResourceProvider struct {
	Options *GenericProviderOpts
}

func NewResourceReaderFromOptions(opt *GenericProviderOpts) (ResourceReader, error) {
	return nil, nil
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