package repoimpl

import (
	"reflect"
	"sync"

	"github.com/hoangduc02011998/tokoin-simple-test/models"
	repo "github.com/hoangduc02011998/tokoin-simple-test/repositories"
)

type UserRepoImpl struct {
	Users *[]models.User
}

type UserSearchWorker struct {
	Search models.Search
	User   models.User
	Users  *[]models.User
}

func NewUserRepo(user *[]models.User) repo.UserRepo {
	return &UserRepoImpl{
		Users: user,
	}
}

// operation process logic
func (job UserSearchWorker) Process() {
	mutex.Lock()
	defer mutex.Unlock()
	value := reflect.ValueOf(job.User)

	modelValue := value.FieldByName(job.Search.FieldName).Interface()
	if isEqualSearch(modelValue, job.Search) {
		*job.Users = append(*job.Users, job.User)
	}
}

// Search with use Worker pool
func (userRepoIml *UserRepoImpl) SearchWithWorker(search models.Search) (*[]models.User, error) {
	queue := make(chan Job)
	killsignal := make(chan bool)
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(*userRepoIml.Users))
	users := &[]models.User{}

	defer func() {
		waitGroup.Wait()
		close(killsignal)
	}()
	// NUMBER_WORKER = 4. so we have 4 thread (goroutine) to process logic
	for i := 0; i < NUMBER_WORKER; i++ {
		go Worker(queue, killsignal, waitGroup, i)
	}
	// assign value UserSearchWorker into channel
	for _, user := range *userRepoIml.Users {
		queue <- UserSearchWorker{
			Search: search,
			User:   user,
			Users:  users,
		}
	}
	return users, nil
}
