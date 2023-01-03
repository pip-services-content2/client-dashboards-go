package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-dashboards-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

type DashboardsClientFixtureV1 struct {
	Client version1.IDashboardsClientV1

	DASHBOARD *version1.DashboardV1
}

func NewDashboardsClientFixtureV1(client version1.IDashboardsClientV1) *DashboardsClientFixtureV1 {
	return &DashboardsClientFixtureV1{
		Client: client,
		DASHBOARD: &version1.DashboardV1{
			UserId: "1",
			App:    "Test App 1",
			Groups: make([]*version1.TileGroupV1, 0),
		},
	}
}

func (c *DashboardsClientFixtureV1) clear() {
	c.Client.DeleteDashboards(context.Background(), "", nil)
}

func (c *DashboardsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	var dashboard1 *version1.DashboardV1

	// Create one dashboard
	dashboard, err := c.Client.GetDashboard(context.Background(), "123", c.DASHBOARD.UserId, c.DASHBOARD.App, "")
	assert.Nil(t, err)

	assert.NotNil(t, dashboard)
	assert.Equal(t, dashboard.UserId, c.DASHBOARD.UserId)
	assert.Equal(t, dashboard.App, c.DASHBOARD.App)

	dashboard1 = dashboard

	// Set the dashboard
	dashboard.Groups = []*version1.TileGroupV1{{Index: 0, Tiles: make([]*version1.TileV1, 0)}}

	dashboard, err = c.Client.SetDashboard(context.Background(), "123", dashboard1)
	assert.Nil(t, err)

	assert.NotNil(t, dashboard)
	assert.Equal(t, dashboard.App, c.DASHBOARD.App)
	assert.Len(t, dashboard.Groups, 1)

	dashboard1 = dashboard

	// Delete dashboard
	err = c.Client.DeleteDashboards(context.Background(), "123", data.NewFilterParamsFromTuples(
		"user_id", c.DASHBOARD.UserId,
		"app", c.DASHBOARD.App,
	))

	assert.Nil(t, err)
}
