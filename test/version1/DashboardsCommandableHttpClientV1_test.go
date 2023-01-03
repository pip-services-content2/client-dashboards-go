package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-dashboards-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type DashboardsCommandableHttpClientV1 struct {
	client  *version1.DashboardsCommandableHttpClientV1
	fixture *DashboardsClientFixtureV1
}

func newDashboardsCommandableHttpClientV1() *DashboardsCommandableHttpClientV1 {
	return &DashboardsCommandableHttpClientV1{}
}

func (c *DashboardsCommandableHttpClientV1) setup(t *testing.T) *DashboardsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewDashboardsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewDashboardsClientFixtureV1(c.client)

	return c.fixture
}

func (c *DashboardsCommandableHttpClientV1) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newDashboardsCommandableHttpClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
