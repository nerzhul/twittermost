package service

import (
	"github.com/labstack/echo/v4"
	"github.com/nerzhul/twittermost/mminteracter/slashcommand"
)

type (
	SlashCommandHandler func(slashcommand.Query, echo.Context) error
)
