package service_organization

import (
	organizationv1 "app/gen/organization/v1"
	"app/model"
	"time"

	"github.com/google/uuid"
)

func FormatterCreateRequestOrganization(req *organizationv1.RequestOrganization) model.Organization {
	return model.Organization{
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		DeletedAt:   nil,
		UUID:        uuid.New().String(),
		Name:        req.Name,
		Domain:      req.Domain,
		Logo:        req.Logo,
		Description: req.Description,
	}
}

func FormatterUpdateRequestOrganization(req *organizationv1.RequestOrganization) model.Organization {
	return model.Organization{
		ID:          int(req.Id),
		UUID:        req.Uuid,
		UpdatedAt:   time.Now().UTC(),
		Name:        req.Name,
		Domain:      req.Domain,
		Logo:        req.Logo,
		Description: req.Description,
	}
}

func FormatterResponseOrganization(org model.Organization) *organizationv1.ResponseOrganization {
	return &organizationv1.ResponseOrganization{
		Uuid:        org.UUID,
		CreatedAt:   org.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   org.UpdatedAt.Format(time.RFC3339),
		Name:        org.Name,
		Domain:      org.Domain,
		Logo:        org.Logo,
		Description: org.Description,
	}
}

func FormatterResponseOrganizationList(orgs []model.Organization) *organizationv1.ResponseOrganizationList {
	var responseList organizationv1.ResponseOrganizationList
	for _, org := range orgs {
		responseList.Organizations = append(responseList.Organizations, FormatterResponseOrganization(org))
	}
	return &responseList
}
