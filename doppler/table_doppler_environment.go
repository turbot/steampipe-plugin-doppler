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
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"Could not find requested environment", "Could not find requested project"}),
			},
			Hydrate: getEnvironment,
		},
		List: &plugin.ListConfig{
			// ParentHydrate: listProjects,
			Hydrate:       listEnvironments,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "project",
					Require: plugin.Optional,
				},
			},
		},
		Columns: commonColumnsForAllResource([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An identifier for the environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "created_at",
				Description: "Date and time of the object's creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "initial_fetch_at",
				Description: "Date and time of the first secrets fetch from a config in the environment.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "project",
				Description: "Identifier of the project the environment belongs to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "slug",
				Description: "A unique identifier for the environment.",
				Type:        proto.ColumnType_STRING,
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listEnvironments(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get client
	client, projectId, err := GetEnvironmentClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_environment.listEnvironments", "client_error", err)
		return nil, err
	}

	input := &doppler.EnvironmentListOptions{
		Project: *projectId,
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

//// HYDRATE FUNCTIONS

func getEnvironment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	slug := d.EqualsQualString("slug")

	// Empty check
	if slug == "" {
		return nil, nil
	}

	// Get client
	client, projectId, err := GetEnvironmentClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_environment.getEnvironment", "client_error", err)
		return nil, err
	}

	input := &doppler.EnvironmentGetOptions{
		Project: *projectId,
		Slug:    slug,
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
