package sample_crm

import (
	crmv1 "app/gen/sample_crm/v1"
	"app/gen/sample_crm/v1/sample_crmv1connect"
	toolsv1 "app/gen/tools/v1"
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type handler struct {
	repo *repository
	sample_crmv1connect.UnimplementedCRMHandler
}

func newHandler(repo *repository) *handler {
	return &handler{repo: repo}
}

func (h *handler) CreateContact(ctx context.Context, req *connect.Request[crmv1.CreateContactRequest]) (*connect.Response[crmv1.Contact], error) {
	contact := CRMContact{
		UUID:    uuid.NewString(),
		Name:    req.Msg.GetName(),
		Email:   req.Msg.GetEmail(),
		Phone:   req.Msg.GetPhone(),
		Company: req.Msg.GetCompany(),
		Notes:   req.Msg.GetNotes(),
	}

	created, err := h.repo.Create(ctx, contact)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(toProtoContact(created)), nil
}

func (h *handler) GetContact(ctx context.Context, req *connect.Request[crmv1.ContactParams]) (*connect.Response[crmv1.Contact], error) {
	contact, err := h.repo.Detail(ctx, req.Msg.GetUuid())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(toProtoContact(contact)), nil
}

func (h *handler) ListContacts(ctx context.Context, req *connect.Request[crmv1.ContactParams]) (*connect.Response[crmv1.ContactListResponse], error) {
	contacts, total, err := h.repo.List(ctx, req.Msg.GetSearch(), int(req.Msg.GetPage()), int(req.Msg.GetPerPage()))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var protoContacts []*crmv1.Contact
	for _, c := range contacts {
		protoContacts = append(protoContacts, toProtoContact(c))
	}

	return connect.NewResponse(&crmv1.ContactListResponse{
		Contacts: protoContacts,
		Total:    int32(total),
	}), nil
}

func (h *handler) UpdateContact(ctx context.Context, req *connect.Request[crmv1.UpdateContactRequest]) (*connect.Response[crmv1.Contact], error) {
	updated, err := h.repo.Update(ctx, req.Msg.GetUuid(), CRMContact{
		Name:    req.Msg.GetName(),
		Email:   req.Msg.GetEmail(),
		Phone:   req.Msg.GetPhone(),
		Company: req.Msg.GetCompany(),
		Notes:   req.Msg.GetNotes(),
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(toProtoContact(updated)), nil
}

func (h *handler) DeleteContact(ctx context.Context, req *connect.Request[crmv1.ContactParams]) (*connect.Response[toolsv1.Empty], error) {
	if err := h.repo.Delete(ctx, req.Msg.GetUuid()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&toolsv1.Empty{}), nil
}

func toProtoContact(c CRMContact) *crmv1.Contact {
	return &crmv1.Contact{
		Id:        int64(c.ID),
		Uuid:      c.UUID,
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		Company:   c.Company,
		Notes:     c.Notes,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
		UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
	}
}
