package businesses

import (
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
	"github.com/hoangduc02011998/tokoin-simple-test/models"
)

type OrganizationView struct {
	Organization   models.Organization
	TicketSubjects []string
	UserNames      []string
}

type OrganizationBiz interface {
	SearchWithWorker(fieldInfo datas.FieldInfo, value string) (*[]models.Organization, error)
	SearchOrganization(fieldInfo datas.FieldInfo, value string) (*[]OrganizationView, error)
}
