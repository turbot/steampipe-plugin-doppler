package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDopplerProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "doppler_project",
		Description: "Doppler Project",
		List: &plugin.ListConfig{
			Hydrate: listProjects,
		},
		Columns: []*plugin.Column{
			{
				Name:        "data",
				Description: "",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromValue(),
			},
		},
	}
}

//// LIST FUNCTION

func listProjects(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// Get client
	svc, err := GetProjectClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_project.listProjects", "client_error", err)
		return nil, err
	}

	op, _, err := svc.List(ctx, &doppler.ProjectListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range op {
		d.StreamLeafListItem(ctx, item)
	}

	return nil, nil
}
