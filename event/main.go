package event

import (
	"github.com/bwmarrin/discordgo"
)

func CreateHandler(s *discordgo.Session) {
	// solely used for diagnostics
	s.AddHandler(MessageCreate)
	s.AddHandler(Ready)

	// user actions
	s.AddHandler(MessageDelete)
	s.AddHandler(MessageUpdate)
	s.AddHandler(GuildMemberAdd)
	s.AddHandler(GuildMemberRemove)

}