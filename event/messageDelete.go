package event

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/misly16/logbot/handler"
)

func MessageDelete(s *discordgo.Session, m * discordgo.MessageDelete) {
	if m.BeforeDelete != nil {
		if m.BeforeDelete.Author.ID == s.State.User.ID {
			return
		}
		embed := &discordgo.MessageEmbed{
			Title: "Message Deleted",
			Description: fmt.Sprintf("Channel: <#%s>\nAuthor: <@%s>", m.ChannelID, m.BeforeDelete.Author.ID),
			Color: 0xe91e63,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Message contents",
					Value:  fmt.Sprintf("```%s```", m.BeforeDelete.Content),
					Inline: false,
				},
			},
		}
		 handler.CreateEmbed(s, embed)
	} else {
	embed := &discordgo.MessageEmbed{
		Title: "Message Deleted",
		Description: fmt.Sprintf("Channel: <#%s>\nMessage not retreivable", m.ChannelID),
		Color: 0xe91e63,
	}
	handler.CreateEmbed(s, embed)
	}
}