package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Golang"
	content := "Content"
	contacts := []string{"email@test", "email2@test"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.NotNil(campaign.CreatedAt)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}
