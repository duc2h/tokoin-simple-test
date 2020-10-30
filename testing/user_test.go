package testing

import (
	"testing"

	bizimpl "github.com/hoangduc02011998/tokoin-simple-test/businesses/business-impl"
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
	"github.com/stretchr/testify/require"
)

func TestSearchUser(t *testing.T) {
	users := datas.NewUser(100000, 3)
	require.Equal(t, 100000-1, len(users))
	inMemory := datas.NewInMemory()
	inMemory.InitInMemory(&users, nil, nil)

	searchable := datas.NewSearchable()
	searchable.InitSearchable()

	if fieldInfo, ok := searchable.FieldUser["_id"]; ok {
		u, _ := bizimpl.NewUserBiz(inMemory, nil).SearchWithWorker(fieldInfo, "89999")

		require.Equal(t, 1, len(*u))
	}

}
