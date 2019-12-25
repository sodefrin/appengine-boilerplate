package controller

import (
	"context"

	"github.com/sodefrin/appengine-boilerplate/proto"
)

// ToHealthzResponse create HealthzResponse.
func ToHealthzResponse(ctx context.Context) *proto.HealthzResponse {
	return &proto.HealthzResponse{}
}
