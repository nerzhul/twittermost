package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nerzhul/twittermost/mminteracter/slashcommand"
)

type Service struct {
	slashCommandRouter map[string]SlashCommandHandler
	e                  *echo.Echo
	port               int
	token              string
}

type HealthcheckStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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

func (s *Service) SetPort(port int) {
	s.port = port
}

func (s *Service) RegisterHealthcheck(path string, h echo.HandlerFunc) {
	s.e.GET(path, h)
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

		// Validate the token
		if len(s.token) != 0 && s.token != cmd.Token {
			return c.String(403, "Invalid token.")
		}

		return h(cmd, c)
	} else {
		return c.JSON(405, nil)
	}
}

func (s *Service) Start() error {
	// Start echo server
	return s.e.Start(fmt.Sprintf(":%d", s.port))
}

func (s *Service) SetAllowedToken(token string) {
	s.token = token
}
