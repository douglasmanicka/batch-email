package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	Id        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

// Initialize  New Campaign
func NewCampaign(name, content string, emails []string) *Campaign {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		Id:   xid.New().String(),
		Name: name,
		//CreatedOn: time.Now(),
		Content:  content,
		Contacts: contacts,
	}

}
