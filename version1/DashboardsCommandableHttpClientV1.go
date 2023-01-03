package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type DashboardsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewDashboardsCommandableHttpClientV1() *DashboardsCommandableHttpClientV1 {
	return &DashboardsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/dashboards"),
	}
}

func (c *DashboardsCommandableHttpClientV1) GetDashboards(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*DashboardV1], error) {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_dashboards", correlationId, params)
	if err != nil {
		return *data.NewEmptyDataPage[*DashboardV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*DashboardV1]](res, correlationId)
}

func (c *DashboardsCommandableHttpClientV1) GetDashboard(ctx context.Context, correlationId string, userId string, app string, kind string) (*DashboardV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"user_id", userId,
		"app", app,
		"kind", kind,
	)

	res, err := c.CallCommand(ctx, "get_dashboard", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*DashboardV1](res, correlationId)
}

func (c *DashboardsCommandableHttpClientV1) SetDashboard(ctx context.Context, correlationId string, dashboard *DashboardV1) (*DashboardV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"dashboard", dashboard,
	)

	res, err := c.CallCommand(ctx, "set_dashboard", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*DashboardV1](res, correlationId)
}

func (c *DashboardsCommandableHttpClientV1) DeleteDashboards(ctx context.Context, correlationId string, filter *data.FilterParams) error {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
	)

	_, err := c.CallCommand(ctx, "delete_dashboards", correlationId, params)
	return err
}
