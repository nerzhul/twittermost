package slashcommand

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Query struct {
	ChannelID   string
	ChannelName string
	Command     string
	ResponseUrl string
	TeamDomain  string
	TeamID      string
	Text        string
	token       string
	UserID      string
	UserName    string
}

func (scq *Query) Deserialize(c echo.Context) error {
	scq.ChannelID = c.FormValue("channel_id")
	scq.ChannelName = c.FormValue("channel_name")
	scq.Command = c.FormValue("command")
	scq.ResponseUrl = c.FormValue("response_url")
	scq.TeamDomain = c.FormValue("team_domain")
	scq.TeamID = c.FormValue("team_id")
	scq.Text = c.FormValue("text")
	scq.UserID = c.FormValue("user_id")
	scq.UserName = c.FormValue("user_name")

	if len(scq.ChannelID) == 0 {
		return errors.New("empty channel_id")
	}

	if len(scq.ChannelName) == 0 {
		return errors.New("empty channel_name")
	}

	if len(scq.Command) == 0 {
		return errors.New("empty command")
	}

	if len(scq.ResponseUrl) == 0 {
		return errors.New("empty response_url")
	}

	if len(scq.TeamDomain) == 0 {
		return errors.New("empty team_domain")
	}

	if len(scq.TeamID) == 0 {
		return errors.New("empty team_id")
	}

	if len(scq.UserID) == 0 {
		return errors.New("empty user_id")
	}

	if len(scq.UserName) == 0 {
		return errors.New("empty user_name")
	}
	return nil
}
