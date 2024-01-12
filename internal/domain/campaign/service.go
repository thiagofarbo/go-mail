package campaign

import (
	"go-mail/internal/contract"
)

type Service struct {
	Repository Repository
	// SendMail   func(campaign *Campaign) error
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	s.Repository.Save(campaign)
	// if err != nil {
	// 	return "", err
	// }
	// err = s.Repository.Create(campaign)
	// if err != nil {
	// 	return "", internalerrors.ErrInternal
	// }

	return campaign.ID, nil
}
