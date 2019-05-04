package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MattermostDialogManager struct {
	slashCommandRouter map[string]echo.HandlerFunc
}

func (m *MattermostDialogManager) Run() {
	m.slashCommandRouter = make(map[string]echo.HandlerFunc)

	// Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	m.registerSlashCommandHandler(e, "/admin", m.handleAdmin)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func (m *MattermostDialogManager) registerSlashCommandHandler(e *echo.Echo, path string, h echo.HandlerFunc) {
	m.slashCommandRouter[path] = h
	e.POST(path, m.handleSlashCommand)
}

func (m *MattermostDialogManager) handleSlashCommand(c echo.Context) error {
	if h, ok := m.slashCommandRouter[c.Path()]; ok {
		return h(c)
	} else {
		return c.JSON(400, nil)
	}
}

func (m *MattermostDialogManager) handleAdmin(c echo.Context) error {
	return c.JSON(400, nil)
}
