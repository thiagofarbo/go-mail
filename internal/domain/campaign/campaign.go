package campaign

import (
	internalerrors "go-mail/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string    `validate: "required"`
	Name      string    `validate: "min=5,max=24"`
	CreatedAt time.Time `validate: "required"`
	Content   string    `validate: "min=5,max=1024"`
	Contacts  []Contact `validate: "min=1"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}
	err := internalerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}
	return nil, err
}
