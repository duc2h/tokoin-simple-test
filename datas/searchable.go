package datas

import (
	"fmt"
	"reflect"

	"github.com/hoangduc02011998/tokoin-simple-test/models"
)

type Searchable struct {
	FieldUser         map[string]FieldInfo
	FieldTicket       map[string]FieldInfo
	FieldOrganization map[string]FieldInfo
}

type FieldInfo struct {
	FielName  string
	FieldType reflect.Kind
}

func NewSearchable() *Searchable {
	return &Searchable{
		FieldUser:         make(map[string]FieldInfo),
		FieldTicket:       make(map[string]FieldInfo),
		FieldOrganization: make(map[string]FieldInfo),
	}
}

// Init Searchable fields for user, ticket, organization
func (searchable *Searchable) InitSearchable() error {

	user := getFieldInfo(models.User{})
	searchable.FieldUser = user

	ticket := getFieldInfo(models.Ticket{})
	searchable.FieldTicket = ticket

	organization := getFieldInfo(models.Organization{})
	searchable.FieldOrganization = organization

	return nil
}

// we get name of field, type of field, tag of field for Searchable
// we use map with key is tag of field and value is FieldInfo
func getFieldInfo(model interface{}) map[string]FieldInfo {
	mapFieldInfo := map[string]FieldInfo{}
	modelType := reflect.TypeOf(model)
	for i := 0; i < modelType.NumField(); i++ {

		field := modelType.Field(i)
		fieldInfo := FieldInfo{
			FielName:  field.Name,
			FieldType: field.Type.Kind(),
		}

		tagName := field.Tag.Get("json")
		mapFieldInfo[tagName] = fieldInfo
	}

	return mapFieldInfo
}

// List of Searchable fields
func (searchable *Searchable) SearchableList() {
	fmt.Println("--------------------------------------")
	fmt.Println("Search Users with: ")
	for k, _ := range searchable.FieldUser {
		fmt.Println(k)
	}
	fmt.Print("\n\n")

	fmt.Println("--------------------------------------")
	fmt.Println("Search Tickets with: ")
	for k, _ := range searchable.FieldTicket {
		fmt.Println(k)
	}
	fmt.Print("\n\n")

	fmt.Println("--------------------------------------")
	fmt.Println("Search Organizations with: ")
	for k, _ := range searchable.FieldOrganization {
		fmt.Println(k)
	}

	fmt.Print("\n\n")
}
