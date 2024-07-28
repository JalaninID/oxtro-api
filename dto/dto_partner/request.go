package dto_partner

import (
	partnerv1 "app/gen/proto/partner/v1"
	"app/model"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

func RequestSwipePartner(req *connect.Request[partnerv1.RequestSwipePartner]) model.Interest {
	return model.Interest{
		ID:             uuid.NewString(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		DeletedAt:      nil,
		IsInterest:     req.Msg.IsInterest,
		InterestUserID: req.Msg.PartnerId,
	}
}
