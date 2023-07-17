package campaign

import (
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

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Test y",
		Content: "Body",
		Emails:  []string{"test@test.com"},
	}

	id, err := service.Create(newCampaign)
	assert.NotEmpty(id)
	assert.Nil(err)

}

func Test_Create_SaveCampaign(t *testing.T) {
	newCampaign := contract.NewCampaign{
		Name:    "Test y",
		Content: "Body",
		Emails:  []string{"test@test.com"},
	}
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

	service := Service{Repository: repositoryMock}
	service.Create(newCampaign)
	repositoryMock.AssertExpectations(t)
}
