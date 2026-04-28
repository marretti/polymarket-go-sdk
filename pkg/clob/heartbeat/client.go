package heartbeat

import (
	"context"

	"github.com/marretti/polymarket-go-sdk/pkg/transport"
)

type Client interface {
	Heartbeat(ctx context.Context, req *HeartbeatRequest) (HeartbeatResponse, error)
}

type clientImpl struct {
	httpClient *transport.Client
}

func NewClient(httpClient *transport.Client) Client {
	return &clientImpl{httpClient: httpClient}
}

func (c *clientImpl) Heartbeat(ctx context.Context, req *HeartbeatRequest) (HeartbeatResponse, error) {
	// Keep-alive: GET /time (server unix time as JSON number or plain body). req is ignored for this route.
	_ = req
	err := c.httpClient.Get(ctx, "/time", nil, nil)
	return HeartbeatResponse{}, err
}
