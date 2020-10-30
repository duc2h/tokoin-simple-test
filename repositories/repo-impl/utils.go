package repoimpl

import (
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/hoangduc02011998/tokoin-simple-test/models"
)

var (
	NUMBER_WORKER = 8
	mutex         = sync.RWMutex{}
)

// Job ...
type Job interface {
	Process()
}

// Worker ...
func Worker(q chan Job, ks chan bool, wg *sync.WaitGroup, numWorker int) {
	for true {
		select {
		case job := <-q:
			func() {
				defer func() {
					if err := recover(); err != nil {
						wg.Done()
					}
				}()

				job.Process()
				wg.Done()
			}()

		case <-ks:
			return
		}
	}
}

// compare two value and return type bool
func isEqualSearch(value interface{}, search models.Search) bool {
	switch search.FieldType {
	case reflect.Int:
		valueConvert := value.(int)
		valueSearch, err := strconv.Atoi(search.Value)
		if err != nil {
			return false
		}
		return valueConvert == valueSearch
	case reflect.String:
		valueConvert := value.(string)
		return valueConvert == search.Value
	case reflect.Bool:
		valueConvert := value.(bool)
		valueSearch, err := strconv.ParseBool(search.Value)
		if err != nil {
			return false
		}

		return valueConvert == valueSearch
	case reflect.Slice:
		typeValue := reflect.TypeOf(value).Elem().Name()

		if typeValue == "string" {
			valueConvert := value.([]string)
			str := strings.Join(valueConvert, ", ")
			return str == search.Value
		}

	}

	return false
}
