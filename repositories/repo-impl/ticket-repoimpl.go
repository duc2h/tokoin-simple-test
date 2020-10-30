package repoimpl

import (
	"reflect"
	"sync"

	"github.com/hoangduc02011998/tokoin-simple-test/models"
	repo "github.com/hoangduc02011998/tokoin-simple-test/repositories"
)

type TicketRepoImpl struct {
	Tickets *[]models.Ticket
}

type TicketSearchWorker struct {
	Search  models.Search
	Ticket  models.Ticket
	Tickets *[]models.Ticket
}

func NewTicketRepo(tickets *[]models.Ticket) repo.TicketRepo {
	return &TicketRepoImpl{
		Tickets: tickets,
	}
}

// operation process logic
func (job TicketSearchWorker) Process() {
	mutex.Lock()
	defer mutex.Unlock()
	value := reflect.ValueOf(job.Ticket)
	modelValue := value.FieldByName(job.Search.FieldName).Interface()

	if isEqualSearch(modelValue, job.Search) {
		*job.Tickets = append(*job.Tickets, job.Ticket)
	}
}

// Search with use Worker pool
func (ticketRepoIml *TicketRepoImpl) SearchWithWorker(search models.Search) (*[]models.Ticket, error) {
	tickets := &[]models.Ticket{}
	queue := make(chan Job)
	killsignal := make(chan bool)
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(*ticketRepoIml.Tickets))
	defer func() {
		waitGroup.Wait()
		close(killsignal)
	}()

	// NUMBER_WORKER = 4. so we have 4 thread to process
	for i := 1; i <= NUMBER_WORKER; i++ {
		go Worker(queue, killsignal, waitGroup, i)
	}

	// assign value TicketSearchWorker into channel
	for _, ticket := range *ticketRepoIml.Tickets {
		queue <- TicketSearchWorker{
			Search:  search,
			Ticket:  ticket,
			Tickets: tickets,
		}
	}

	return tickets, nil
}
