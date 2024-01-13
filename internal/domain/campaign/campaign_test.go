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

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.NotNil(campaign.CreatedAt)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Golang"
	content := "Content"
	contacts := []string{"email@test", "email2@test"}

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)
	name := ""
	content := "Content"
	contacts := []string{"email@test", "email2@test"}
	_, err := NewCampaign(name, content, contacts)

	assert.Equal("Name can not be empty or null", err.Error())
}
