package datas

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hoangduc02011998/tokoin-simple-test/models"
)

func init() {
	// to change value when random
	rand.Seed(time.Now().UnixNano())
}

func NewOrganization(n int) []models.Organization {
	organizations := []models.Organization{}

	for i := 1; i < n; i++ {
		org := models.Organization{
			Id:            i,
			Url:           GetUrl(fmt.Sprint("organizations/", i, ".json")),
			ExternalId:    randomID(),
			Name:          strings.Title(randomName(4)),
			DomainNames:   randomDomain(4),
			CreatedAt:     randomDateTime(),
			Details:       strings.Title(randomName(8)),
			SharedTickets: randomBool(),
			Tags:          randomTag(4),
		}

		organizations = append(organizations, org)
	}

	return organizations
}

func NewUser(n int, orgN int) []models.User {
	users := []models.User{}

	for i := 1; i < n; i++ {
		firstName := randomName(3)
		lastName := randomName(4)
		fullName := firstName + " " + lastName
		alias := randomName(4)

		user := models.User{
			Id:             i,
			Url:            GetUrl(fmt.Sprint("users/", i, ".json")),
			ExternalId:     randomID(),
			Name:           strings.Title(fullName),
			Alias:          fmt.Sprint(randomGender(), " ", alias),
			CreatedAt:      randomDateTime(),
			Active:         randomBool(),
			Verified:       randomBool(),
			Shared:         randomBool(),
			Locale:         randomLocale(),
			Timezone:       randomTimezone(),
			LastLoginAt:    randomDateTime(),
			Email:          randomEmail(alias, lastName),
			Phone:          RandomPhone(),
			Signature:      "Don't Worry Be Happy!",
			OrganizationId: rand.Intn(orgN-1) + 1,
			Tags:           randomTag(4),
			Suspended:      true,
			Role:           randomRole(),
		}

		users = append(users, user)
	}

	return users
}

func NewTicket(n int, orgN int, userN int) []models.Ticket {
	tickets := []models.Ticket{}

	for i := 1; i < n; i++ {
		id := randomID()
		ticket := models.Ticket{
			Id:             id,
			Url:            GetUrl(fmt.Sprint("tickets/", id, ".json")),
			ExternalId:     randomID(),
			CreatedAt:      randomDateTime(),
			Type:           RandomType(),
			Subject:        randomText(3),
			Description:    randomText(10),
			Priority:       randomPriority(),
			Status:         randomStatus(),
			SubmitterId:    rand.Intn(userN-1) + 1,
			AssigneeId:     rand.Intn(userN-1) + 1,
			OrganizationId: rand.Intn(orgN-1) + 1,
			Tags:           randomTag(4),
			HasIncidents:   randomBool(),
			DueAt:          randomDateTime(),
			Via:            "chat",
		}

		tickets = append(tickets, ticket)
	}

	return tickets
}
