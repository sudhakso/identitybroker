package mapper

import (
	"errors"
	
	provider "github.com/identitybroker/pkg/provider"
	api "github.com/identitybroker/api/_generated"
)

type ResourceSerializer interface {
	Serialize(res *provider.ProviderResource) (*api.Resource, error)
	SerializeCollection(res *provider.ProviderResources) (*api.Resources, error)
	// Types supported
	SerializeType(t *provider.ProviderResourceType) (*api.ResourceType, error)
}

type ResourceMapper struct {
	serializer map[string]ResourceSerializer
}

func (r* ResourceMapper) Serialize(res *provider.ProviderResource) (*api.Resource, error) {
	// delegate
	if result, err := r.serializer[res.Kind].Serialize(res); err == nil{
		return result, nil
	}
	return nil, nil
}

func (r* ResourceMapper) SerializeCollection(res *provider.ProviderResources) (*api.Resources, error) {
	// delegate
	results := api.Resources{Resources:make([]*api.Resource, 10, 20)}
	if res != nil {
		vals := res.Resources
		for _,val := range vals {
			sval, err := r.serializer[val.Kind].Serialize(val)
			if err != nil {
				results.Resources = append(results.Resources, sval)
			}
		}
	} else {
		return nil, errors.New("Empty resource")
	}
	return &results, nil
}

func (r* ResourceMapper) SerializeType(t *provider.ProviderResourceType) (*api.ResourceType, error) {
	return nil, nil
}

func NewResourceMapper() (*ResourceMapper){
	serializer := make(map[string]ResourceSerializer)
	
	//Add all serializers we have
	serializer["user"] = &UserTypeConversion{}
	serializer["app"]  = &ApplicationTypeConversion{}
	serializer["type"] = &ResourceTypeConversion{}
	
	return &ResourceMapper{serializer: serializer}
}

type UserTypeConversion struct {	
}

func (u* UserTypeConversion) Serialize(res *provider.ProviderResource) (*api.Resource, error) {
	// TBD: Serialize User
	if res != nil {
		r := &api.Resource{Kind: "User", Name: res.Name, Id: res.Id}
		return r, nil
	}
	return nil, errors.New("Serialization Failed for ProviderResource")
}


func (u* UserTypeConversion) SerializeType(t *provider.ProviderResourceType) (*api.ResourceType, error) {
	// TBD: Serialize UserType
	if t != nil {
		r := &api.ResourceType{Name: t.Name, PathPrefix: t.Location}
		return r, nil
	}
	return nil, errors.New("Serialization Failed for ProviderResourceType")
}

func (u* UserTypeConversion) SerializeCollection(res *provider.ProviderResources) (*api.Resources, error) {
	if res != nil {
		r := api.Resources{Resources: make([]*api.Resource, 1, 2)}
		ut := &UserTypeConversion{}
		for _, pr := range res.Resources {
			_pr, err :=  ut.Serialize(pr)
			if err != nil {
				return nil, errors.New("Serialization Failed for ProviderResources")
			} else {
				r.Resources = append(r.Resources, _pr)
			}
		}
		return &r, nil
	}
	return nil, errors.New("Serialization Failed for ProviderResources")
}

type ApplicationTypeConversion struct {	
}

func (u* ApplicationTypeConversion) Serialize(res *provider.ProviderResource) (*api.Resource, error) {
	// TBD: Serialize Application
	return nil, nil
}

func (u* ApplicationTypeConversion) SerializeType(t *provider.ProviderResourceType) (*api.ResourceType, error) {
	// TBD: Serialize UserType
	return nil, nil
}

func (u* ApplicationTypeConversion) SerializeCollection(res *provider.ProviderResources) (*api.Resources, error) {
	return nil, nil
}

type ResourceTypeConversion struct {	
}

func (u* ResourceTypeConversion) Serialize(res *provider.ProviderResource) (*api.Resource, error) {
	// NOP
	return nil, nil
}

func (u* ResourceTypeConversion) SerializeType(t *provider.ProviderResourceType) (*api.ResourceType, error) {
	// TBD: Serialize UserType
	return nil, nil
}

func (u* ResourceTypeConversion) SerializeCollection(res *provider.ProviderResources) (*api.Resources, error) {
	return nil, nil
}

