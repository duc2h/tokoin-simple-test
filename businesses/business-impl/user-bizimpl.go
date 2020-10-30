package businessesimpl

import (
	"strconv"

	biz "github.com/hoangduc02011998/tokoin-simple-test/businesses"
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
	"github.com/hoangduc02011998/tokoin-simple-test/models"
	repo "github.com/hoangduc02011998/tokoin-simple-test/repositories"
	repoimpl "github.com/hoangduc02011998/tokoin-simple-test/repositories/repo-impl"
)

type UserBizImpl struct {
	UserRepo   repo.UserRepo
	memory     *datas.InMemory
	searchable *datas.Searchable
}

func NewUserBiz(memory *datas.InMemory, searchable *datas.Searchable) biz.UserBiz {
	return &UserBizImpl{
		UserRepo:   repoimpl.NewUserRepo(memory.Users),
		memory:     memory,
		searchable: searchable,
	}
}

func (userBizImpl *UserBizImpl) SearchWithWorker(fieldInfo datas.FieldInfo, value string) (*[]models.User, error) {
	search := models.Search{
		FieldName: fieldInfo.FielName,
		FieldType: fieldInfo.FieldType,
		Value:     value,
	}

	users, err := userBizImpl.UserRepo.SearchWithWorker(search)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Show all info of User, all of relationships
func (userBizImpl *UserBizImpl) SearchUser(fieldInfo datas.FieldInfo, value string) (*[]biz.UserView, error) {
	userViews := []biz.UserView{}

	users, err := userBizImpl.SearchWithWorker(fieldInfo, value)
	if err != nil {
		return nil, err
	}

	ticketBiz := NewTicketBiz(userBizImpl.memory, userBizImpl.searchable)
	orgBiz := NewOrganizationBiz(userBizImpl.memory, userBizImpl.searchable)

	for _, user := range *users {

		// subject of tickets by submitterId
		fieldSubmitter, _ := userBizImpl.searchable.FieldTicket["submitter_id"]
		submitters := getTicketSubjects(ticketBiz, fieldSubmitter, strconv.Itoa(user.Id))

		// subject of tickets by assigneeId
		fieldAssignee, _ := userBizImpl.searchable.FieldTicket["assignee_id"]
		assignees := getTicketSubjects(ticketBiz, fieldAssignee, strconv.Itoa(user.Id))

		// get organization name
		orgIdField, _ := userBizImpl.searchable.FieldOrganization["_id"]
		orgName := getOrganizationName(orgBiz, orgIdField, strconv.Itoa(user.OrganizationId))

		userView := biz.UserView{
			Users:            user,
			TicketSubmitters: submitters,
			TicketAssignees:  assignees,
			OrganizationName: orgName,
		}

		userViews = append(userViews, userView)
	}

	return &userViews, nil
}
