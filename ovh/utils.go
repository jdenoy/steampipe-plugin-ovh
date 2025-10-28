package ovh

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ovh/go-ovh/ovh"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*ovh.Client, error) {
	// get ovh client from cache
	cacheKey := fmt.Sprintf("ovh_%s", d.Connection.Name)
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*ovh.Client), nil
	}

	applicationKey := ""
	applicationSecret := ""
	consumerKey := ""
	endpoint := ""

	ovhConfig := GetConfig(d.Connection)

	// Configuration with precedence: config file > environment variables
	if ovhConfig.ApplicationKey != nil {
		applicationKey = *ovhConfig.ApplicationKey
	} else if key := os.Getenv("OVH_APPLICATION_KEY"); key != "" {
		applicationKey = key
	}

	if ovhConfig.ApplicationSecret != nil {
		applicationSecret = *ovhConfig.ApplicationSecret
	} else if secret := os.Getenv("OVH_APPLICATION_SECRET"); secret != "" {
		applicationSecret = secret
	}

	if ovhConfig.ConsumerKey != nil {
		consumerKey = *ovhConfig.ConsumerKey
	} else if key := os.Getenv("OVH_CONSUMER_KEY"); key != "" {
		consumerKey = key
	}

	if ovhConfig.Endpoint != nil {
		endpoint = *ovhConfig.Endpoint
	} else if ep := os.Getenv("OVH_ENDPOINT"); ep != "" {
		endpoint = ep
	}

	if applicationKey == "" {
		return nil, errors.New("'application_key' must be set in the connection configuration or OVH_APPLICATION_KEY environment variable. Edit your connection configuration file and then restart Steampipe")
	}
	if applicationSecret == "" {
		return nil, errors.New("'application_secret' must be set in the connection configuration or OVH_APPLICATION_SECRET environment variable. Edit your connection configuration file and then restart Steampipe")
	}
	if consumerKey == "" {
		return nil, errors.New("'consumer_key' must be set in the connection configuration or OVH_CONSUMER_KEY environment variable. Edit your connection configuration file and then restart Steampipe")
	}
	if endpoint == "" {
		return nil, errors.New("'endpoint' must be set in the connection configuration or OVH_ENDPOINT environment variable. Edit your connection configuration file and then restart Steampipe")
	}

	client, err := ovh.NewClient(
		endpoint,
		applicationKey,
		applicationSecret,
		consumerKey,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OVH client: %w", err)
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

// ValidateQualValue validates that a required qualifier has a non-empty string value.
func ValidateQualValue(qualName, value string) error {
	if value == "" {
		return fmt.Errorf("'%s' qualifier must be provided in the where clause", qualName)
	}
	return nil
}
