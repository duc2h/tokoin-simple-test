package testing

import (
	"testing"

	bizimpl "github.com/hoangduc02011998/tokoin-simple-test/businesses/business-impl"
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
	"github.com/stretchr/testify/require"
)

func TestSearchTicket(t *testing.T) {
	tickets := datas.NewTicket(1000, 200, 10)
	require.Equal(t, 1000-1, len(tickets))
	inMemory := datas.NewInMemory()
	inMemory.InitInMemory(nil, &tickets, nil)

	searchable := datas.NewSearchable()
	searchable.InitSearchable()

	if fieldInfo, ok := searchable.FieldTicket["via"]; ok {
		u, _ := bizimpl.NewTicketBiz(inMemory, nil).SearchWithWorker(fieldInfo, "chat")

		require.Equal(t, 1000-1, len(*u))
	}

}
