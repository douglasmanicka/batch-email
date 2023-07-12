package campaign

import "github.com/douglasmanicka/batch-email/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) error {
	return nil
}
