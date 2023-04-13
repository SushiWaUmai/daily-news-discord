package commands

import "github.com/bwmarrin/discordgo"

func init() {
	createCommand(&discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Sends pong",
	}, func(dg *discordgo.Session, i *discordgo.InteractionCreate) {
		dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData {
				Content: "pong!",
			},
		})
	})
}
