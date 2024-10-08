package Command

import (
	"errors"

	"github.com/bwmarrin/discordgo"

	"github.com/shibaisdog/opns/Command/Slash"
	"github.com/shibaisdog/opns/Error"
)

type Slash_Definition *discordgo.ApplicationCommand

type Setup_Slash struct {
	Definition Slash_Definition
	Handler    func(Slash.Event)
}

var Slash_CommandList = []Setup_Slash{}

// Register the slash_command
func (S *Setup_Slash) Register() {
	if S.Definition == nil || S.Handler == nil {
		Error.New(Error.Err{
			Msg: errors.New("warning: Slash command is nil"),
		}, true)
		return
	}
	Slash_CommandList = append(Slash_CommandList, *S)
}
