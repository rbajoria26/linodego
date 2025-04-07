package linodego

import (
	"context"
)

// MonitorServiceToken represents a MonitorServiceToken object
type MonitorServiceToken struct {
	Token string `json:"token"`
}

// Create token options
type MonitorTokenCreateOptions struct {
	EntityIds []int `json:"entity_ids"`
}

// ListMonitorServiceTokenByServiceType to create token for a given serviceType
func (c *Client) CreateMonitorServiceTokenForServiceType(ctx context.Context, serviceType string, opts MonitorTokenCreateOptions) (*MonitorServiceToken, error) {
	e := formatAPIPath("monitor/services/%s/token", serviceType)
	return doPOSTRequest[MonitorServiceToken](ctx, c, e, opts)
}
