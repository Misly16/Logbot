package event

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, m * discordgo.Ready) {
	fmt.Println("Ready to GO!")
	
}