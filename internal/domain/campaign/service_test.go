package campaign

import (
	"errors"
	"testing"

	"github.com/douglasmanicka/batch-email/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test y",
		Content: "Body",
		Emails:  []string{"test@test.com"},
	}
	newCampaign2 = contract.NewCampaign{
		Name:    "Test z",
		Content: "Body2",
		Emails:  []string{"test2@test.com"},
	}

	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)
	assert.NotEmpty(id)
	//assert.NotNil(id)
	assert.Nil(err)

}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	//for quick test use below
	// repositoryMock.On("Save", mock.Anything).Return(nil)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	//service := Service{Repository: repositoryMock}
	service.Repository = repositoryMock
	service.Create(newCampaign)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_Campaign_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""
	_, err := service.Create(newCampaign)
	assert.NotNil(err)
	assert.Equal("name is required", err.Error())

}

func Test_Create_ValidateRepositorySaveCampaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	//for quick test use below
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock

	_, err := service.Create(newCampaign2)
	assert.Equal("error to save on database", err.Error())

}
