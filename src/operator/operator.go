package operator

import (
	"net"

	"github.com/sr/operator/src/gcloud"
	"github.com/sr/operator/src/papertrail"
	"go.pedge.io/env"
	"google.golang.org/grpc"
)

func Listen() error {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}
	server := grpc.NewServer()

	gcloud.RegisterGCloudServiceServer(server, gcloud.NewAPIServer())

	papertrailEnv := &papertrail.Env{}
	if err := env.Populate(papertrailEnv); err != nil {
		return err
	}
	papertrail.RegisterPapertrailServiceServer(
		server,
		papertrail.NewAPIServer(papertrailEnv),
	)

	return server.Serve(listener)
}
