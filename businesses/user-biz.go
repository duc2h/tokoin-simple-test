package businesses

import (
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
	"github.com/hoangduc02011998/tokoin-simple-test/models"
)

type UserView struct {
	Users            models.User
	TicketSubmitters []string
	TicketAssignees  []string
	OrganizationName string
}

type UserBiz interface {
	SearchWithWorker(fieldInfo datas.FieldInfo, value string) (*[]models.User, error)
	SearchUser(fieldInfo datas.FieldInfo, value string) (*[]UserView, error)
}
