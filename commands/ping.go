package commands

import "github.com/bwmarrin/discordgo"

func init() {
	var defaultMemberPermissions int64 = discordgo.PermissionViewChannel | discordgo.PermissionSendMessages
	dmPermission := true

	createCommand(&discordgo.ApplicationCommand{
		Name:                     "ping",
		Description:              "Sends pong",
		DefaultMemberPermissions: &defaultMemberPermissions,
		DMPermission:             &dmPermission,
	}, func(dg *discordgo.Session, i *discordgo.InteractionCreate) {
		dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "pong!",
			},
		})
	})
}
