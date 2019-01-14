package provider

import (

)

type ProviderConnection struct { //ProviderOpts
  dummy string
}

type RegisterOpts struct { //ProviderRegistrationOpts
	dummy string
}

type Status struct { //RegistrationStatus
	dummy string
}

type Registrar interface {
	RegisterProvider() (*Status, error) 
	DeRegisterProvider() (*Status, error)
} 

// TBD
type ProviderRegistrationFactory struct {
	Options *RegisterOpts
}

func NewProviderRegistrar(opts *RegisterOpts) (*Registrar, error) {
	return nil, nil
}

func (r *ProviderRegistrationFactory) RegisterProvider() (*Status, error) {
	return nil, nil
}

func (r *ProviderRegistrationFactory) DeRegisterProvider() (*Status, error) {
	return nil, nil
}