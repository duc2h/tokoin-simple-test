package testing

import (
	"fmt"
	"testing"

	bizimpl "github.com/hoangduc02011998/tokoin-simple-test/businesses/business-impl"
	"github.com/hoangduc02011998/tokoin-simple-test/datas"
	"github.com/stretchr/testify/require"
)

func TestSearchOrg(t *testing.T) {
	orgs := datas.NewOrganization(1000)
	require.Equal(t, 1000-1, len(orgs))
	inMemory := datas.NewInMemory()
	inMemory.InitInMemory(nil, nil, &orgs)

	searchable := datas.NewSearchable()
	searchable.InitSearchable()

	if fieldInfo, ok := searchable.FieldOrganization["url"]; ok {
		u, _ := bizimpl.NewOrganizationBiz(inMemory, nil).SearchWithWorker(fieldInfo, datas.GetUrl(fmt.Sprint("organizations/", 10, ".json")))

		require.Equal(t, 1, len(*u))
	}

}
