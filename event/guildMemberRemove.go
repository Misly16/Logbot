package event

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/misly16/logbot/handler"
)

func GuildMemberRemove(s *discordgo.Session, m * discordgo.GuildMemberRemove) {
	embed := &discordgo.MessageEmbed{
		Title: "Member Left",
		Description: fmt.Sprintf("<@%s>", m.Member.User.ID),
		Color: 0xe91e63,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Joined server",
				Value: fmt.Sprintf("<t:%d> **(<t:%d:R>)**", m.Member.JoinedAt.Unix(), m.Member.JoinedAt.Unix()),
				Inline: false,
			},
		},
	}
	 handler.CreateEmbed(s, embed)
}