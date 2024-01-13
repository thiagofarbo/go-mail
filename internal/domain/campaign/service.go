package campaign

import (
	"go-mail/internal/contract"
	internalerrors "go-mail/internal/internal-errors"
)

type Service struct {
	Repository Repository
	// SendMail   func(campaign *Campaign) error
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}
