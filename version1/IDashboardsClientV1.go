package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IDashboardsClientV1 interface {
	GetDashboards(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*DashboardV1], error)

	GetDashboard(ctx context.Context, correlationId string, userId string, app string, kind string) (*DashboardV1, error)

	SetDashboard(ctx context.Context, correlationId string, dashboard *DashboardV1) (*DashboardV1, error)

	DeleteDashboards(ctx context.Context, correlationId string, filter *data.FilterParams) error
}
