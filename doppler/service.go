package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	"github.com/nikoksr/doppler-go/project"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func GetProjectClient(ctx context.Context, d *plugin.Connection) (*project.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		return &project.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.APIKey,
		}, nil
	}

	return nil, nil
}
