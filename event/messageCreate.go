package event

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m * discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "log!ping" {
		pingStart := time.Now()
		message, _ := s.ChannelMessageSendReply(m.ChannelID, "pong", m.Reference())
		pingEnd := time.Since(pingStart)
		s.ChannelMessageEdit(m.ChannelID, message.ID, "pong - "+pingEnd.String())

	}
}