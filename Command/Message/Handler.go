package Message

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fluffy-melli/opns/Traffic"
)

type Event struct {
	Traffic     *Traffic.Count
	Interaction *discordgo.MessageCreate
	Client      *discordgo.Session
}

type Response_Message struct {
	Message *discordgo.Message
	Handler *Event
}
