package repositories

import "github.com/hoangduc02011998/tokoin-simple-test/models"

type TicketRepo interface {
	SearchWithWorker(search models.Search) (*[]models.Ticket, error)
}
