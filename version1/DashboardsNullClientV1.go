package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type DashboardsNullClientV1 struct {
}

func NewDashboardsNullClientV1() *DashboardsNullClientV1 {
	return &DashboardsNullClientV1{}
}

func (c *DashboardsNullClientV1) GetDashboards(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*DashboardV1], error) {
	return *data.NewEmptyDataPage[*DashboardV1](), nil
}

func (c *DashboardsNullClientV1) GetDashboard(ctx context.Context, correlationId string, userId string, app string, kind string) (*DashboardV1, error) {
	return nil, nil
}

func (c *DashboardsNullClientV1) SetDashboard(ctx context.Context, correlationId string, dashboard *DashboardV1) (*DashboardV1, error) {
	return dashboard, nil
}

func (c *DashboardsNullClientV1) DeleteDashboards(ctx context.Context, correlationId string, filter *data.FilterParams) error {
	panic("not implemented") // TODO: Implement
}
