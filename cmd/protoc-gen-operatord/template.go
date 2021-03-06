package main

import "github.com/sr/operator/generator"

var (
	mainTemplate = generator.NewTemplate("main-gen.go",
		`// Code generated by protoc-gen-operatord
package main

import (
	"os"
	"fmt"

	"github.com/sr/operator"
	"google.golang.org/grpc"
	"go.pedge.io/env"

{{range .Services}}
	"{{.ImportPath}}"
{{end}}
)

func run() error {
	config, err := operator.NewConfigFromEnv()
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	logger := operator.NewLogger()
	instrumenter := operator.NewInstrumenter(logger)
	server := operator.NewServer(grpcServer, config, logger, instrumenter)
{{range .Services}}
	{{.Name}}Env := &{{.PackageName}}.Env{}
	if err := env.Populate({{.Name}}Env); err != nil {
		server.LogServiceStartupError("{{.Name}}", err)
	} else {
		if {{.Name}}Server, err := {{.PackageName}}.NewAPIServer({{.Name}}Env); err != nil {
			server.LogServiceStartupError("{{.Name}}", err)
		} else {
			instrumented := &instrumented{{.PackageName}}{{.FullName}}{instrumenter, {{.Name}}Server}
			{{.Name}}.Register{{camelCase .FullName}}Server(grpcServer, instrumented)
			server.LogServiceRegistered("{{.Name}}")
		}
	}
{{end}}
	return server.Serve()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "operatord: %s\n", err)
		os.Exit(1)
	}
}
`)

	instrumentedTemplate = generator.NewTemplate("instrumented-gen.go",
		`// Code generated by protoc-gen-operatord
package main

import (
	"time"

	"github.com/sr/operator"
	"golang.org/x/net/context"

	servicepkg "{{.ImportPath}}"
)

type instrumented{{.PackageName}}{{.FullName}} struct {
	instrumenter operator.Instrumenter
	server       servicepkg.{{.FullName}}Server
}

{{- range .Methods}}
// {{.Name}} instruments the {{$.FullName}}.{{.Name}} method.
func (a *instrumented{{$.PackageName}}{{$.FullName}}) {{.Name}}(
	ctx context.Context,
	request *servicepkg.{{.Input}},
) (response *servicepkg.{{.Output}}, err error) {
	defer func(start time.Time) {
		a.instrumenter.Instrument(
			operator.NewRequest(
				request.Source,
				"{{$.PackageName}}",
				"{{.Name}}",
				"{{.Input}}",
				"{{.Output}}",
				err,
				start,
			),
		)
	}(time.Now())
	return a.server.{{.Name}}(ctx, request)
}
{{end}}`)
)
