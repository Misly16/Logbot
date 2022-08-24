package event

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/misly16/logbot/handler"
)

func MessageUpdate(s *discordgo.Session, m * discordgo.MessageUpdate) {
	if m.BeforeUpdate != nil {
		if m.BeforeUpdate.Author.ID == s.State.User.ID {
			return
		}
		embed := &discordgo.MessageEmbed{
			Title: "Message Edited",
			Color: 0xf1c40f,
			Description: fmt.Sprintf("Channel: <#%s>\nAuthor: <@%s>", m.ChannelID, m.BeforeUpdate.Author.ID),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Before",
					Value:  fmt.Sprintf("```%s```", m.BeforeUpdate.Content),
					Inline: false,
				},
				{
					Name:   "After",
					Value:  fmt.Sprintf("```%s```", m.Content),
					Inline: false,
				},
			},
		}
		handler.CreateEmbed(s, embed)
	} else {
	embed := &discordgo.MessageEmbed{
		Title: "Message Edited",
		Description: fmt.Sprintf("Channel: <#%s>", m.ChannelID),
		Color: 0xf1c40f,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Before",
				Value:  "```Message not retreivable```",
				Inline: false,
			},
			{
				Name:   "After",
				Value:  fmt.Sprintf("```%s```", m.Content),
				Inline: false,
			},
		},
	}
	handler.CreateEmbed(s, embed)
	}
}