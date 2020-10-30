package businessesimpl

import (
	biz "github.com/hoangduc02011998/tokoin-simple-test/businesses"
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
)

func getUserNames(userBiz biz.UserBiz, fieldInfo datas.FieldInfo, value string) []string {
	userNames := []string{}
	users, err := userBiz.SearchWithWorker(fieldInfo, value)
	if err != nil || len(*users) == 0 {
		return userNames
	}

	for _, user := range *users {
		userNames = append(userNames, user.Name)
	}

	return userNames
}

func getTicketSubjects(ticketBiz biz.TicketBiz, fieldInfo datas.FieldInfo, value string) []string {
	subjects := []string{}
	tickets, err := ticketBiz.SearchWithWorker(fieldInfo, value)

	if err != nil {
		return nil
	}

	for _, ticket := range *tickets {
		subjects = append(subjects, ticket.Subject)
	}

	return subjects
}

func getOrganizationName(orgBiz biz.OrganizationBiz, fieldInfo datas.FieldInfo, value string) string {
	orgs, err := orgBiz.SearchWithWorker(fieldInfo, value)

	if err != nil || len(*orgs) == 0 {
		return ""
	}

	organizations := *orgs
	return organizations[0].Name
}
