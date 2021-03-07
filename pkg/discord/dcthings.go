package discord

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordThings struct {
	DiscordToken string
	DiscordSession *discordgo.Session

}


func NewDiscordThings(token string) (*DiscordThings, error){
	dc := new(DiscordThings)
	dc.DiscordToken = token
	var err error
	dc.DiscordSession, err = discordgo.New("Bot " + dc.DiscordToken)
	return dc, err
}



func (d *DiscordThings) TestMessage(s *discordgo.Session, m *discordgo.MessageCreate){
	d.DiscordSession.ChannelMessageSend(m.ChannelID, "Weapico")
}