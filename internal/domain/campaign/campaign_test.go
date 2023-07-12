package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"a@gmail.com", "b@gmail.com"}
)

func TestNewCampaign(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}
func TestNewCampaign_IdIsNotEmpty(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.NotEmpty(campaign.Id)
}
func TestNewCampaign_CreateOnMustBeNow(t *testing.T) {
	//Arrange
	assert := assert.New(t)
	nowLessOne := time.Now().Add(-time.Minute)

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.Greater(campaign.CreatedOn, nowLessOne)
}

func TestNewCampaign_MustValidateName(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Act
	_, error := NewCampaign("", content, contacts)

	//Assert
	assert.Equal("name is required", error.Error())
}

func TestNewCampaign_MustValidateContent(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Act
	_, error := NewCampaign(name, "", contacts)

	//Assert
	assert.Equal("content is required", error.Error())
}

func TestNewCampaign_MustValidateContacts(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Act
	_, error := NewCampaign(name, content, []string{})

	//Assert
	assert.Equal("contacts is required", error.Error())
}
