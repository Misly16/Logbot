package handler

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func CreateEmbed(s *discordgo.Session, embed *discordgo.MessageEmbed) {
	s.ChannelMessageSendEmbed(os.Getenv("DISCORD_LOG_CHANNEL"), embed)
}