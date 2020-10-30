package repositories

import "github.com/hoangduc02011998/tokoin-simple-test/models"

type OrganizationRepo interface {
	SearchWithWorker(search models.Search) (*[]models.Organization, error)
}
