package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	"github.com/nikoksr/doppler-go/audit"
	projectConfig "github.com/nikoksr/doppler-go/config"
	"github.com/nikoksr/doppler-go/environment"
	"github.com/nikoksr/doppler-go/project"
	"github.com/nikoksr/doppler-go/secret"
	servicetoken "github.com/nikoksr/doppler-go/service_token"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func GetProjectClient(ctx context.Context, d *plugin.Connection) (*project.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		doppler.Key = *config.APIKey
		return &project.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, nil
	}

	return nil, nil
}

func GetConfigClient(ctx context.Context, d *plugin.Connection) (*projectConfig.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		doppler.Key = *config.APIKey
		return &projectConfig.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, nil
	}

	return nil, nil
}

func GetSecretClient(ctx context.Context, d *plugin.Connection) (*secret.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		doppler.Key = *config.APIKey
		return &secret.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, nil
	}

	return nil, nil
}

func GetEnvironmentClient(ctx context.Context, d *plugin.Connection) (*environment.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		return &environment.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.APIKey,
		}, nil
	}

	return nil, nil
}

func GetServiceTokenClient(ctx context.Context, d *plugin.Connection) (*servicetoken.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		return &servicetoken.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.APIKey,
		}, nil
	}

	return nil, nil
}

func GetUserClient(ctx context.Context, d *plugin.Connection) (*audit.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		return &audit.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.APIKey,
		}, nil
	}

	return nil, nil
}
