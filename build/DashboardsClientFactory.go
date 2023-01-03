package build

import (
	clients1 "github.com/pip-services-content2/client-dashboards-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type DashboardsClientFactory struct {
	*cbuild.Factory
}

func NewDashboardsClientFactory() *DashboardsClientFactory {
	c := &DashboardsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-dashboards", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-dashboards", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-dashboards", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewDashboardsNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewDashboardsMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewDashboardsCommandableHttpClientV1)

	return c
}
