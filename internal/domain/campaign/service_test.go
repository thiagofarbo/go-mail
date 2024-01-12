package campaign

import (
	"go-mail/internal/contract"
	"testing"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (repository *repositoryMock) Save(campaign *Campaign) error {
	args := repository.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {

	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test@mail.com"},
	}

	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {

		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)

}
