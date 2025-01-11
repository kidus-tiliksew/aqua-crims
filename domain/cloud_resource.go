package domain

import "github.com/stackus/errors"

var (
	ErrCloudResIDEmpty         = errors.Wrap(errors.ErrBadRequest, "id cannot be blank")
	ErrCloudResNameEmpty       = errors.Wrap(errors.ErrBadRequest, "name cannot be blank")
	ErrCloudResTypeEmpty       = errors.Wrap(errors.ErrBadRequest, "type cannot be blank")
	ErrCloudResRegionEmpty     = errors.Wrap(errors.ErrBadRequest, "region cannot be blank")
	ErrCloudResCustomerIDEmpty = errors.Wrap(errors.ErrBadRequest, "customer id cannot be blank")
)

type CloudResource struct {
	ID         int64
	Name       string
	Type       string
	Region     string
	CustomerID int64
}

func CreateCloudResource(name, resourceType, region string, customerID int64) (*CloudResource, error) {
	if name == "" {
		return nil, ErrCloudResNameEmpty
	}

	if resourceType == "" {
		return nil, ErrCloudResTypeEmpty
	}

	if region == "" {
		return nil, ErrCloudResRegionEmpty
	}

	return &CloudResource{
		Name:       name,
		Type:       resourceType,
		Region:     region,
		CustomerID: customerID,
	}, nil
}

func UpdateCloudResource(id int64, customerID int64, name, resourceType, region string) (*CloudResource, error) {
	if id == 0 {
		return nil, ErrCloudResIDEmpty
	}

	if customerID == 0 {
		return nil, ErrCloudResCustomerIDEmpty
	}

	if name == "" {
		return nil, ErrCloudResNameEmpty
	}

	if resourceType == "" {
		return nil, ErrCloudResTypeEmpty
	}

	if region == "" {
		return nil, ErrCloudResRegionEmpty
	}

	return &CloudResource{
		ID:         id,
		Name:       name,
		Type:       resourceType,
		Region:     region,
		CustomerID: customerID,
	}, nil
}
