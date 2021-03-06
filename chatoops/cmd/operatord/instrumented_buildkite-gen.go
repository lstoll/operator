// Code generated by protoc-gen-operatord
package main

import (
	"time"

	"github.com/sr/operator"
	"golang.org/x/net/context"

	servicepkg "github.com/sr/operator/chatoops/services/buildkite"
)

type instrumentedbuildkiteBuildkiteService struct {
	instrumenter operator.Instrumenter
	server       servicepkg.BuildkiteServiceServer
}

// Status instruments the BuildkiteService.Status method.
func (a *instrumentedbuildkiteBuildkiteService) Status(
	ctx context.Context,
	request *servicepkg.StatusRequest,
) (response *servicepkg.StatusResponse, err error) {
	defer func(start time.Time) {
		a.instrumenter.Instrument(
			operator.NewRequest(
				request.Source,
				"buildkite",
				"Status",
				"StatusRequest",
				"StatusResponse",
				err,
				start,
			),
		)
	}(time.Now())
	return a.server.Status(ctx, request)
}

// ListBuilds instruments the BuildkiteService.ListBuilds method.
func (a *instrumentedbuildkiteBuildkiteService) ListBuilds(
	ctx context.Context,
	request *servicepkg.ListBuildsRequest,
) (response *servicepkg.ListBuildsResponse, err error) {
	defer func(start time.Time) {
		a.instrumenter.Instrument(
			operator.NewRequest(
				request.Source,
				"buildkite",
				"ListBuilds",
				"ListBuildsRequest",
				"ListBuildsResponse",
				err,
				start,
			),
		)
	}(time.Now())
	return a.server.ListBuilds(ctx, request)
}
