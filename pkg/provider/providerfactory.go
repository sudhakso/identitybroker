package provider

import (

)

type ProviderConnection struct { //ProviderOpts
  Dummy string
}

type RegisterOpts struct { //ProviderRegistrationOpts
	Dummy string
}

type Status struct { //RegistrationStatus
	Dummy string
}

type Registrar interface {
	RegisterProvider() (*Status, error) 
	DeRegisterProvider() (*Status, error)
} 

// TBD
type ProviderRegistrationFactory struct {
	Options *RegisterOpts
}

func NewProviderRegistrar(opts *RegisterOpts) (Registrar, error) {
	return nil, nil
}

func (r *ProviderRegistrationFactory) RegisterProvider() (*Status, error) {
	return nil, nil
}

func (r *ProviderRegistrationFactory) DeRegisterProvider() (*Status, error) {
	return nil, nil
}