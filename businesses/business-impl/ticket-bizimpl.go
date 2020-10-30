package businessesimpl

import (
	"strconv"

	biz "github.com/hoangduc02011998/tokoin-simple-test/businesses"
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
	"github.com/hoangduc02011998/tokoin-simple-test/models"
	repo "github.com/hoangduc02011998/tokoin-simple-test/repositories"
	repoimpl "github.com/hoangduc02011998/tokoin-simple-test/repositories/repo-impl"
)

type TicketBizImpl struct {
	TicketRepo repo.TicketRepo
	memory     *datas.InMemory
	searchable *datas.Searchable
}

func NewTicketBiz(memory *datas.InMemory, searchable *datas.Searchable) biz.TicketBiz {
	return &TicketBizImpl{
		TicketRepo: repoimpl.NewTicketRepo(memory.Tickets),
		memory:     memory,
		searchable: searchable,
	}
}

func (ticketBizImpl *TicketBizImpl) SearchWithWorker(fieldInfo datas.FieldInfo, value string) (*[]models.Ticket, error) {
	search := models.Search{
		FieldName: fieldInfo.FielName,
		FieldType: fieldInfo.FieldType,
		Value:     value,
	}

	tickets, err := ticketBizImpl.TicketRepo.SearchWithWorker(search)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (ticketBizImpl *TicketBizImpl) SearchTicket(fieldInfo datas.FieldInfo, value string) (*[]biz.TicketView, error) {
	ticketViews := []biz.TicketView{}
	tickets, err := ticketBizImpl.SearchWithWorker(fieldInfo, value)
	if err != nil {
		return nil, err
	}

	userBiz := NewUserBiz(ticketBizImpl.memory, ticketBizImpl.searchable)
	orgBiz := NewOrganizationBiz(ticketBizImpl.memory, ticketBizImpl.searchable)

	for _, ticket := range *tickets {
		// get assigneeName and submiiterName from user
		idField, _ := ticketBizImpl.searchable.FieldUser["_id"]
		assigneeName := getUserNames(userBiz, idField, strconv.Itoa(ticket.AssigneeId))[0]
		submitterName := getUserNames(userBiz, idField, strconv.Itoa(ticket.SubmitterId))[0]

		// get organizationName from Organization
		orgIdField, _ := ticketBizImpl.searchable.FieldOrganization["_id"]
		orgName := getOrganizationName(orgBiz, orgIdField, strconv.Itoa(ticket.OrganizationId))

		ticketView := biz.TicketView{
			Ticket:           ticket,
			AssingeeName:     assigneeName,
			SubmitterName:    submitterName,
			OrganizationName: orgName,
		}

		ticketViews = append(ticketViews, ticketView)
	}

	return &ticketViews, nil
}
