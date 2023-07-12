package campaign

import (
	"testing"

	"github.com/douglasmanicka/batch-email/internal/contract"
	"github.com/stretchr/testify/assert"
)

func Test_Create_campaign(t *testing.T) {
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
