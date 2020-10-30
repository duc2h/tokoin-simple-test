package businesses

import (
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
	"github.com/hoangduc02011998/tokoin-simple-test/models"
)

type TicketView struct {
	Ticket           models.Ticket
	AssingeeName     string
	SubmitterName    string
	OrganizationName string
}

type TicketBiz interface {
	SearchWithWorker(fieldInfo datas.FieldInfo, value string) (*[]models.Ticket, error)
	SearchTicket(fieldInfo datas.FieldInfo, value string) (*[]TicketView, error)
}
