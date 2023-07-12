package domain

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
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}
func TestNewCampaign_IdIsNotEmpty(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.NotEmpty(campaign.Id)
}
func TestNewCampaign_CreateOnMustBeNow(t *testing.T) {
	//Arrange
	assert := assert.New(t)
	nowLessOne := time.Now().Add(-time.Minute)

	//Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Greater(campaign.CreatedOn, nowLessOne)
}
