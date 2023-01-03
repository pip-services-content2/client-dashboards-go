package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type DashboardsMockClientV1 struct {
	dashboards []*DashboardV1
}

func NewDashboardsMockClientV1() *DashboardsMockClientV1 {
	return &DashboardsMockClientV1{
		dashboards: make([]*DashboardV1, 0),
	}
}

func (c *DashboardsMockClientV1) composeFilter(filter *data.FilterParams) func(item *DashboardV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	id, idOk := filter.GetAsNullableString("id")
	userId, userIdOk := filter.GetAsNullableString("user_id")
	app, appOk := filter.GetAsNullableString("app")
	kind, kindOk := filter.GetAsNullableString("kind")

	return func(item *DashboardV1) bool {
		if idOk && id != item.Id {
			return false
		}
		if userIdOk && userId != item.UserId {
			return false
		}
		if appOk && app != item.App {
			return false
		}
		if kindOk && kind != item.Kind {
			return false
		}

		return true
	}
}

func (c *DashboardsMockClientV1) GetDashboards(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*DashboardV1], error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*DashboardV1, 0)
	for _, v := range c.dashboards {
		item := *v
		if filterFunc(&item) {
			item.Id = "" // delete id
			items = append(items, &item)
		}
	}
	return *data.NewDataPage(items, len(c.dashboards)), nil
}

func (c *DashboardsMockClientV1) makeDashboardId(userId string, app string, kind string) string {
	id := userId + "_" + app
	if kind != "" {
		id += "_" + kind
	}
	return id
}

func (c *DashboardsMockClientV1) GetDashboard(ctx context.Context, correlationId string, userId string, app string, kind string) (*DashboardV1, error) {
	id := c.makeDashboardId(userId, app, kind)
	var dashboard *DashboardV1

	for _, ds := range c.dashboards {
		if ds.Id == id {
			dashboard = ds
			break
		}
	}

	if dashboard == nil {
		dashboard = &DashboardV1{
			UserId: userId,
			App:    app,
			Kind:   kind,
		}
	}

	if dashboard != nil {
		dashboard.Id = ""
	}

	return dashboard, nil
}

func (c *DashboardsMockClientV1) SetDashboard(ctx context.Context, correlationId string, dashboard *DashboardV1) (*DashboardV1, error) {
	dashboard.Id = c.makeDashboardId(dashboard.UserId, dashboard.App, dashboard.Kind)

	var index *int

	for i, ds := range c.dashboards {
		if ds.Id == dashboard.Id {
			index = &i
			break
		}
	}

	buf := *dashboard

	if index != nil {
		c.dashboards[*index] = &buf
	} else {
		c.dashboards = append(c.dashboards, &buf)
	}

	dashboard.Id = ""

	return dashboard, nil
}

func (c *DashboardsMockClientV1) DeleteDashboards(ctx context.Context, correlationId string, filter *data.FilterParams) error {
	filterFunc := c.composeFilter(filter)

	for index, v := range c.dashboards {
		item := *v
		if filterFunc(&item) {
			c.dashboards = append(c.dashboards[:index], c.dashboards[index+1:]...)
		}
	}
	return nil
}
