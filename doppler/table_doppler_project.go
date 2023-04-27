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
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getProject,
		},
		List: &plugin.ListConfig{
			Hydrate: listProjects,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "ID is the unique identifier for the object.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time when the project was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "The description of the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "slug",
				Description: "Slug is an abbreviated name for the project.",
				Type:        proto.ColumnType_STRING,
			},

			// Doppler standard column
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Get client
	client, err := GetProjectClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_project.listProjects", "client_error", err)
		return nil, err
	}

	input := &doppler.ProjectListOptions{}

	// The SDK does not support pagination till date(04/23).
	op, _, err := client.List(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_project.listProjects", "api_error", err)
		return nil, err
	}

	for _, item := range op {
		d.StreamListItem(ctx, item)

		// Context may get cancelled due to manual cancellation or if the limit has been reached.
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	projectName := d.EqualsQualString("name")

	// Get client
	client, err := GetProjectClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_project.getProject", "client_error", err)
		return nil, err
	}

	input := &doppler.ProjectGetOptions{
		Name: projectName,
	}

	op, _, err := client.Get(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_project.getProject", "api_error", err)
		return nil, err
	}

	return op, nil
}
