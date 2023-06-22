package doppler

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDopplerProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "doppler_project",
		Description: "A Doppler project defines app config and secrets for one service or app.",
		List: &plugin.ListConfig{
			Hydrate: getProject,
			// If the .spc file has an Invalid project id (the project id is not available in the workplace). In that case, if we run the query 'select * from doppler_project', it will throw a not found error.
			// Since we make the API call for this table using the project ID value specified in the .spc file, ignoring the not found error for this table will result in an empty row being returned. However, this behavior is generally not reported to the user, leaving them unaware of why the table is returning no results.
			// So not found error has not been ingnored here.
		},
		Columns: commonColumnsForAllResource([]*plugin.Column{
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

func getProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// A Project ID field is mandatory in the connection configuration file, and it must be present in each connection.
	// For each connection, we are fetching project details from the cache.
	project, err := getProjectData(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_project.getProject", "getProjectData_cached", err)
		return nil, err
	}
	d.StreamListItem(ctx, project)

	return nil, nil
}
