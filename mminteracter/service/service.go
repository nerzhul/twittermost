package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nerzhul/twittermost/mminteracter/slashcommand"
)

type Service struct {
	slashCommandRouter map[string]SlashCommandHandler
	e                  *echo.Echo
}

func New() *Service {
	s := &Service{}
	// Echo instance
	s.e = echo.New()
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.slashCommandRouter = make(map[string]SlashCommandHandler)
	return s
}

func (s *Service) RegisterSlashCommandHandler(path string, h SlashCommandHandler) {
	s.slashCommandRouter[path] = h
	s.e.POST(path, s.handleSlashCommand)
}

func (s *Service) handleSlashCommand(c echo.Context) error {
	if h, ok := s.slashCommandRouter[c.Path()]; ok {
		var cmd slashcommand.Query
		if err := cmd.Deserialize(c); err != nil {
			return c.JSON(400, err)
		}
		// @TODO: verify if we have all we need in the request
		return h(cmd, c)
	} else {
		return c.JSON(405, nil)
	}
}

func (s *Service) Start() error {
	// Start echo server
	return s.e.Start(":8080")
}
