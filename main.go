package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	bizimpl "github.com/hoangduc02011998/tokoin-simple-test/businesses/business-impl"
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
)

func main() {
	// init data on memory
	inMemory := datas.NewInMemory()
	err := inMemory.InitDataFromFile()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// init searchable
	searchable := datas.NewSearchable()
	err = searchable.InitSearchable()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Println("Select search options:")
		fmt.Println(". Press 1 to search ")
		fmt.Println(". Press 2 to view a list of searchable fields")
		fmt.Println(". Type 'quit' to exit")
		var option string
		if scanner.Scan() {
			option = scanner.Text()
		}
		// fmt.Scan(&option)

		switch option {
		case "quit":
			return
		case "1":
			fmt.Println("Select 1) Users or 2) Tickets or 3) Organizations")
			var optionSearch string
			if scanner.Scan() {
				optionSearch = scanner.Text()
			}

			fmt.Println("Enter search term")
			var term string
			if scanner.Scan() {
				term = scanner.Text()
			}

			fmt.Println("Enter search value")
			var value string
			if scanner.Scan() {
				value = scanner.Text()
			}

			if optionSearch == "1" {
				if fieldInfo, ok := searchable.FieldUser[term]; ok {
					users, err := bizimpl.NewUserBiz(inMemory, searchable).SearchUser(fieldInfo, value)
					if err != nil || len(*users) == 0 {
						fmt.Println("No results found")
						break
					} else {
						showData(*users)
					}

				} else {
					fmt.Println("No results found")
				}
			} else if optionSearch == "2" {
				if fieldInfo, ok := searchable.FieldTicket[term]; ok {
					tickets, err := bizimpl.NewTicketBiz(inMemory, searchable).SearchTicket(fieldInfo, value)
					if err != nil || len(*tickets) == 0 {
						fmt.Println("No results found")
						break
					} else {
						showData(*tickets)
					}

				} else {
					fmt.Println("No results found")
				}
			} else if optionSearch == "3" {
				if fieldInfo, ok := searchable.FieldOrganization[term]; ok {
					orgs, err := bizimpl.NewOrganizationBiz(inMemory, searchable).SearchOrganization(fieldInfo, value)
					if err != nil || len(*orgs) == 0 {
						fmt.Println("No results found")
						break
					} else {
						showData(*orgs)
					}

				} else {
					fmt.Println("No results found")
				}
			}

		case "2":
			searchable.SearchableList()
		}

	}
}

func showData(datas interface{}) {
	b, _ := json.MarshalIndent(datas, " ", "    ")
	fmt.Println(string(b))
}
