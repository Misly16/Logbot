package event

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/misly16/logbot/handler"
)

func GuildMemberAdd(s *discordgo.Session, m * discordgo.GuildMemberAdd) {
	int64ID, _ := strconv.ParseInt(m.Member.User.ID, 10, 64)
	embed := &discordgo.MessageEmbed{
		Title: "Member Joined",
		Description: fmt.Sprintf("<@%s>", m.Member.User.ID),
		Color: 0x2ecc71,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Account created",
				Value: fmt.Sprintf("<t:%d> **(<t:%d:R>)**", handler.CreationTime(int64ID), handler.CreationTime(int64ID)),
				Inline: false,
			},
		},
	}
	 handler.CreateEmbed(s, embed)
}