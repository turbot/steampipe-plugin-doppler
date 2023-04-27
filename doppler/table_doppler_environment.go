package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDopplerEnvironment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "doppler_environment",
		Description: "Doppler Environment",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project", "slug"}),
			Hydrate:    getEnvironment,
		},
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listEnvironments,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "project",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "An identifier for the environment.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "slug",
				Type:        proto.ColumnType_STRING,
				Description: "A unique identifier for the environment.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the environment.",
			},
			{
				Name:        "project",
				Type:        proto.ColumnType_STRING,
				Description: "Identifier of the project the environment belongs to.",
			},
			{
				Name:        "initial_fetch_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Date and time of the first secrets fetch from a config in the environment.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Date and time of the object's creation.",
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

func listEnvironments(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*doppler.Project)
	projectId := d.EqualsQualString("project")

	// Reduce the numbers of API call if the project id is provided in the where clause.
	if projectId != "" && projectId != *project.ID {
		return nil, nil
	}

	// Get client
	client, err := GetEnvironmentClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_environment.listEnvironments", "client_error", err)
		return nil, err
	}

	input := &doppler.EnvironmentListOptions{
		Project: *project.ID,
	}

	// The SDK does not support pagination till date(04/23).
	op, _, err := client.List(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_environment.listEnvironments", "api_error", err)
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

func getEnvironment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	projectId := d.EqualsQualString("project")
	slugId := d.EqualsQualString("slug")

	// Get client
	client, err := GetEnvironmentClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_environment.getEnvironment", "client_error", err)
		return nil, err
	}

	input := &doppler.EnvironmentGetOptions{
		Project: projectId,
		Slug:    slugId,
	}

	op, _, err := client.Get(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_environment.getEnvironment", "api_error", err)
		return nil, err
	}

	if op != nil {
		return op, nil
	}

	return nil, nil
}
