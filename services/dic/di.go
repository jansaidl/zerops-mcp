package dic

import (
	"github.com/jansaidl/zerops-mcp/services/server"
	"github.com/jansaidl/zerops-mcp/tools/container"
	"github.com/zerops-dev/di/appRunner"
	"github.com/zerops-dev/di/di/s"
	_ "golang.org/x/tools/imports"
)

//go:generate templater -tags "di" -templateTags "!templater,!di"
var _ = func() *s.Di {
	di := s.NewDi("zerops-mcp",
		s.WithCommand(
			"run",
			"run command",
			``,
		),
	)

	di.Add(
		s.Scope(
			appRunner.DiScope(),
			s.Service(server.New, s.WithSetter(server.RegisterTools)),

			s.Service(container.NewReadDir),
			s.Service(container.NewReadFile),
			s.Service(container.NewWriteFile),

			s.Config(server.NewConfig, "server"),
		),
	)

	return di
}
