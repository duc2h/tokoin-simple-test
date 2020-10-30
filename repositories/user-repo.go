package repositories

import "github.com/hoangduc02011998/tokoin-simple-test/models"

type UserRepo interface {
	SearchWithWorker(search models.Search) (*[]models.User, error)
}
