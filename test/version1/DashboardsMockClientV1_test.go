package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-dashboards-go/version1"
)

type DashboardsMockClientV1 struct {
	client  *version1.DashboardsMockClientV1
	fixture *DashboardsClientFixtureV1
}

func newDashboardsMockClientV1() *DashboardsMockClientV1 {
	return &DashboardsMockClientV1{}
}

func (c *DashboardsMockClientV1) setup(t *testing.T) *DashboardsClientFixtureV1 {
	c.client = version1.NewDashboardsMockClientV1()
	c.fixture = NewDashboardsClientFixtureV1(c.client)
	return c.fixture
}

func (c *DashboardsMockClientV1) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newDashboardsMockClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
