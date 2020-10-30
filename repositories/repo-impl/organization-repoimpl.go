package repoimpl

import (
	"reflect"
	"sync"

	"github.com/hoangduc02011998/tokoin-simple-test/models"
	repo "github.com/hoangduc02011998/tokoin-simple-test/repositories"
)

type OrganizationRepoImpl struct {
	Organizations *[]models.Organization
}

type OrganizationSearchWorker struct {
	Search        models.Search
	Organization  models.Organization
	Organizations *[]models.Organization
}

func NewOrganization(organizations *[]models.Organization) repo.OrganizationRepo {
	return &OrganizationRepoImpl{
		Organizations: organizations,
	}
}

func (job OrganizationSearchWorker) Process() {
	mutex.Lock()
	defer mutex.Unlock()
	value := reflect.ValueOf(job.Organization)
	modelValue := value.FieldByName(job.Search.FieldName).Interface()

	if isEqualSearch(modelValue, job.Search) {
		*job.Organizations = append(*job.Organizations, job.Organization)
	}
}

// Search with use Worker pool
func (orgRepoImpl *OrganizationRepoImpl) SearchWithWorker(search models.Search) (*[]models.Organization, error) {
	organizations := &[]models.Organization{}
	queue := make(chan Job)
	killsignal := make(chan bool)
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(*orgRepoImpl.Organizations))
	defer func() {
		waitGroup.Wait()
		close(killsignal)
	}()

	// NUMBER_WORKER = 4. so we have 4 thread to process
	for i := 1; i <= NUMBER_WORKER; i++ {
		go Worker(queue, killsignal, waitGroup, i)
	}

	// assign value OrganizationRepoImpl into channel
	for _, org := range *orgRepoImpl.Organizations {
		queue <- OrganizationSearchWorker{
			Search:        search,
			Organization:  org,
			Organizations: organizations,
		}
	}

	return organizations, nil
}
