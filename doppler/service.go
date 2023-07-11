package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	activitylog "github.com/nikoksr/doppler-go/activity_log"
	"github.com/nikoksr/doppler-go/audit"
	projectConfig "github.com/nikoksr/doppler-go/config"
	"github.com/nikoksr/doppler-go/environment"
	"github.com/nikoksr/doppler-go/project"
	"github.com/nikoksr/doppler-go/secret"
	servicetoken "github.com/nikoksr/doppler-go/service_token"
	"github.com/nikoksr/doppler-go/workplace"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Project Client
func GetProjectClient(ctx context.Context, d *plugin.Connection) (*project.Client, error) {
	config := GetConfigWithToken(d)

	doppler.Key = *config.Token
	return &project.Client{
		Backend: doppler.GetBackend(),
		Key:     doppler.Key,
	}, nil

}

// Config Client
func GetConfigClient(ctx context.Context, d *plugin.Connection) (*projectConfig.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.Token != "" {
		doppler.Key = *config.Token
		return &projectConfig.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, config.ProjectId, nil
	}

	return nil, nil, nil
}

// Secret Client
func GetSecretClient(ctx context.Context, d *plugin.Connection) (*secret.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.Token != "" || *config.ProjectId != "" {
		doppler.Key = *config.Token
		return &secret.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, config.ProjectId, nil
	}

	return nil, nil, nil
}

// Environment Client
func GetEnvironmentClient(ctx context.Context, d *plugin.Connection) (*environment.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.Token != "" || *config.ProjectId != "" {
		return &environment.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.Token,
		}, config.ProjectId, nil
	}

	return nil, nil, nil
}

// Service Token Client
func GetServiceTokenClient(ctx context.Context, d *plugin.Connection) (*servicetoken.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.Token != "" || *config.ProjectId != "" {
		return &servicetoken.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.Token,
		}, config.ProjectId, nil
	}

	return nil, nil, nil
}

// User Client
func GetUserClient(ctx context.Context, d *plugin.Connection) (*audit.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.Token != "" || *config.ProjectId != "" {
		return &audit.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.Token,
		}, config.ProjectId, nil
	}
	return nil, nil, nil
}

// Workplace Client
func GetWorkplaceClient(ctx context.Context, d *plugin.Connection) (*workplace.Client, error) {
	config := GetConfigWithToken(d)

	if *config.Token != "" || *config.ProjectId != "" {
		return &workplace.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.Token,
		}, nil
	}

	return nil, nil
}

func GetActivityLogClient(ctx context.Context, d *plugin.Connection) (*activitylog.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.Token != "" || *config.ProjectId != "" {
		return &activitylog.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.Token,
		}, config.ProjectId, nil
	}

	return nil, nil, nil
}
