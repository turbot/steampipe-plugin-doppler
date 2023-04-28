package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const (
	// Constants for Standard Column Descriptions
	ColumnDescriptionWorkplaceName = "Title of the resource."
	ColumnDescriptionWorkplaceId   = "The ID of the workplace."
	ColumnDescriptionTitle         = "The name of the workplace."
)

func commonColumnsForAllResource(columns []*plugin.Column) []*plugin.Column {
	return append(columns, []*plugin.Column{
		{
			Name:        "workplace_name",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: ColumnDescriptionWorkplaceName,
			Transform:   transform.FromField("Name"),
		},
		{
			Name:        "workplace_id",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: ColumnDescriptionWorkplaceId,
			Transform:   transform.FromField("ID"),
		},
	}...)
}

// returns the workplace_id, workplace_name common column which is added across all the tables
func getCommonColumns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// Trace logging to debug cache and execution flows
	plugin.Logger(ctx).Debug("getCommonColumns", "status", "starting", "connection_name", d.Connection.Name)

	var workplaceData *doppler.Workplace

	workplace, err := getWorkplace(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("getCommonColumns", "status", "failed", "connection_name", d.Connection.Name, "error", err)
		return nil, err
	}

	workplaceData = workplace.(*doppler.Workplace)

	plugin.Logger(ctx).Debug("getCommonColumns", "status", "finished", "connection_name", d.Connection.Name, "workplaceData", workplaceData)

	return workplaceData, nil
}

var getWorkplace = plugin.HydrateFunc(getWorkplaceDataUncached).Memoize()

func getWorkplaceDataUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create Client
	client, err := GetWorkplaceClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("getWorkplaceDataUncached", "status", "failed", "connection_name", d.Connection.Name, "client_error", err)
		return nil, err
	}

	response, _, err := client.Get(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("getWorkplaceDataUncached", "status", "failed", "connection_name", d.Connection.Name, "api_error", err)
		return nil, err
	}

	return response, nil
}

// Get the project details
func getProjectData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	var projectData *doppler.Project

	project, err := getProjectDetails(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("getCommonColumns", "status", "failed", "connection_name", d.Connection.Name, "error", err)
		return nil, err
	}

	projectData = project.(*doppler.Project)
	plugin.Logger(ctx).Error("Project name", *projectData.ID)

	return projectData, nil

}

var getProjectDetails = plugin.HydrateFunc(getProjectDataUncached).Memoize()

func getProjectDataUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// get Connection config
	config := GetConfigWithToken(d.Connection)

	// Create Client
	client, err := GetProjectClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("getProjectDataUncached", "status", "failed", "connection_name", d.Connection.Name, "client_error", err)
		return nil, err
	}

	response, _, err := client.Get(ctx, &doppler.ProjectGetOptions{
		Name: *config.PROJECT_ID,
	})
	if err != nil {
		plugin.Logger(ctx).Error("getProjectDataUncached", "status", "failed", "connection_name", d.Connection.Name, "api_error", err)
		return nil, err
	}

	return response, nil
}
