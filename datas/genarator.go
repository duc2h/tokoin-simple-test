package datas

import (
	"encoding/json"
	"io/ioutil"
	"reflect"

	"github.com/hoangduc02011998/tokoin-simple-test/models"
)

type InMemory struct {
	Users         *[]models.User
	Tickets       *[]models.Ticket
	Organizations *[]models.Organization
}

func NewInMemory() *InMemory {
	return &InMemory{
		Users:         &[]models.User{},
		Tickets:       &[]models.Ticket{},
		Organizations: &[]models.Organization{},
	}
}

func (inMemory *InMemory) InitInMemory(Users *[]models.User, tickets *[]models.Ticket, orgs *[]models.Organization) {
	inMemory.Users = Users
	inMemory.Tickets = tickets
	inMemory.Organizations = orgs
}

// import data from file json
func (inMemory *InMemory) InitDataFromFile() error {
	// read file user.json and assign value from file into Users InMemory
	dataUsers, err := readFile("./datas/users.json", []models.User{})
	if err != nil {
		return err
	}
	inMemory.Users = dataUsers.(*[]models.User)

	// read file ticket and assign value from file into Tickets InMemory
	dataTickets, err := readFile("./datas/tickets.json", []models.Ticket{})
	if err != nil {
		return err
	}
	inMemory.Tickets = dataTickets.(*[]models.Ticket)

	// read file organization and assign value from file into Organizations InMemory
	dataOrganizations, err := readFile("./datas/organizations.json", []models.Organization{})
	if err != nil {
		return err
	}
	inMemory.Organizations = dataOrganizations.(*[]models.Organization)

	return nil
}

func readFile(fileName string, model interface{}) (interface{}, error) {
	dataByte, err := ioutil.ReadFile(fileName)
	modelType := reflect.TypeOf(model)
	modelValue := reflect.New(reflect.SliceOf(modelType).Elem()).Interface()
	if err != nil {
		return nil, err
	}

	json.Unmarshal(dataByte, &modelValue)

	return modelValue, nil
}
