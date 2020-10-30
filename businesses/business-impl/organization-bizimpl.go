package businessesimpl

import (
	"strconv"

	biz "github.com/hoangduc02011998/tokoin-simple-test/businesses"
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
	"github.com/hoangduc02011998/tokoin-simple-test/models"
	repo "github.com/hoangduc02011998/tokoin-simple-test/repositories"
	repoimpl "github.com/hoangduc02011998/tokoin-simple-test/repositories/repo-impl"
)

type OrganizationBizImpl struct {
	OrganizationRepo repo.OrganizationRepo
	memory           *datas.InMemory
	searchable       *datas.Searchable
}

func NewOrganizationBiz(memory *datas.InMemory, searchable *datas.Searchable) biz.OrganizationBiz {
	return &OrganizationBizImpl{
		OrganizationRepo: repoimpl.NewOrganization(memory.Organizations),
		memory:           memory,
		searchable:       searchable,
	}
}

func (orgBizImpl *OrganizationBizImpl) SearchWithWorker(fieldInfo datas.FieldInfo, value string) (*[]models.Organization, error) {
	search := models.Search{
		FieldName: fieldInfo.FielName,
		FieldType: fieldInfo.FieldType,
		Value:     value,
	}

	organizations, err := orgBizImpl.OrganizationRepo.SearchWithWorker(search)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func (orgBizImpl *OrganizationBizImpl) SearchOrganization(fieldInfo datas.FieldInfo, value string) (*[]biz.OrganizationView, error) {
	orgViews := []biz.OrganizationView{}
	organizations, err := orgBizImpl.SearchWithWorker(fieldInfo, value)

	if err != nil {
		return nil, err
	}
	userBiz := NewUserBiz(orgBizImpl.memory, orgBizImpl.searchable)
	ticketBiz := NewTicketBiz(orgBizImpl.memory, orgBizImpl.searchable)
	for _, org := range *organizations {

		orgIdField, _ := orgBizImpl.searchable.FieldUser["organization_id"]

		userNames := getUserNames(userBiz, orgIdField, strconv.Itoa(org.Id))
		ticketSubjects := getTicketSubjects(ticketBiz, orgIdField, strconv.Itoa(org.Id))

		orgView := biz.OrganizationView{
			Organization:   org,
			UserNames:      userNames,
			TicketSubjects: ticketSubjects,
		}

		orgViews = append(orgViews, orgView)
	}

	return &orgViews, nil
}
