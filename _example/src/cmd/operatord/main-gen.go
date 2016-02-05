// Code generated by protoc-gen-operatord
package main

import (
	"os"
	"fmt"

	"services/buildkite"

	"services/controller"

	"services/gcloud"

	"services/papertrail"

	"google.golang.org/grpc"
	"github.com/sr/operator/server"
	"go.pedge.io/env"
)

func run() error {
	config, err := server.NewConfigFromEnv()
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	logger := server.NewLogger()
	instrumentator := server.NewInstrumentator(logger)
	server := server.NewServer(grpcServer, config, logger, instrumentator)

	buildkiteEnv := &buildkite.Env{}
	if err := env.Populate(buildkiteEnv); err != nil {
		server.LogServiceStartupError("buildkite", err)
	} else {
		if buildkiteServer, err := buildkite.NewAPIServer(buildkiteEnv); err != nil {
			server.LogServiceStartupError("buildkite", err)
		} else {
			instrumented := buildkite.NewInstrumentedBuildkiteServiceServer(instrumentator, buildkiteServer)
			buildkite.RegisterBuildkiteServiceServer(grpcServer, instrumented)
			server.LogServiceRegistered("buildkite")
		}
	}

	controllerEnv := &controller.Env{}
	if err := env.Populate(controllerEnv); err != nil {
		server.LogServiceStartupError("controller", err)
	} else {
		if controllerServer, err := controller.NewAPIServer(controllerEnv); err != nil {
			server.LogServiceStartupError("controller", err)
		} else {
			instrumented := controller.NewInstrumentedControllerServer(instrumentator, controllerServer)
			controller.RegisterControllerServer(grpcServer, instrumented)
			server.LogServiceRegistered("controller")
		}
	}

	gcloudEnv := &gcloud.Env{}
	if err := env.Populate(gcloudEnv); err != nil {
		server.LogServiceStartupError("gcloud", err)
	} else {
		if gcloudServer, err := gcloud.NewAPIServer(gcloudEnv); err != nil {
			server.LogServiceStartupError("gcloud", err)
		} else {
			instrumented := gcloud.NewInstrumentedGcloudServiceServer(instrumentator, gcloudServer)
			gcloud.RegisterGcloudServiceServer(grpcServer, instrumented)
			server.LogServiceRegistered("gcloud")
		}
	}

	papertrailEnv := &papertrail.Env{}
	if err := env.Populate(papertrailEnv); err != nil {
		server.LogServiceStartupError("papertrail", err)
	} else {
		if papertrailServer, err := papertrail.NewAPIServer(papertrailEnv); err != nil {
			server.LogServiceStartupError("papertrail", err)
		} else {
			instrumented := papertrail.NewInstrumentedPapertrailServiceServer(instrumentator, papertrailServer)
			papertrail.RegisterPapertrailServiceServer(grpcServer, instrumented)
			server.LogServiceRegistered("papertrail")
		}
	}

	return server.Serve()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "operatord: %s\n", err)
		os.Exit(1)
	}
}
