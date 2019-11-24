package eventvalidating

type Service interface {
	ValidateLatestHistoryEvent(shipmentID string) error
}

type Repository interface {
	LatestHistoryEventIsPublished(shipmentID string) error
}

type service struct {
	validationRepo Repository
}

func NewService(validationRepo Repository) Service {
	return &service{validationRepo}
}

func (srv *service) ValidateLatestHistoryEvent(shipmentID string) error {
	return srv.validationRepo.LatestHistoryEventIsPublished(shipmentID)
}
