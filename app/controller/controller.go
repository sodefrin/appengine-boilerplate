package controller

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sodefrin/appengine-boilerplate/log"
	"github.com/sodefrin/appengine-boilerplate/proto"
)

// AppengineBoilerplate is struct for grpc server.
type AppengineBoilerplate struct {
}

// NewAppengineBoilerplate create new AppengineBoilerplate struct.
func NewAppengineBoilerplate() *AppengineBoilerplate {
	return &AppengineBoilerplate{}
}

// Healthz is health check endpoint.
func (c *AppengineBoilerplate) Healthz(ctx context.Context, _ *empty.Empty) (*proto.HealthzResponse, error) {
	log.Logger.Info("started to Healthz")

	return ToHealthzResponse(ctx), nil
}
