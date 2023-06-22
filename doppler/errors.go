package doppler

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// shouldIgnoreErrors:: function which returns an ErrorPredicate for Doppler API calls
func shouldIgnoreErrors(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {

		for _, pattern := range notFoundErrors {
			if ok := strings.Contains(err.Error(), pattern); ok {
				return true
			}
		}

		return false
	}
}