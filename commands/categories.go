package commands

import (
	"strings"

	"github.com/SushiWaUmai/daily-news-discord/news"
	"github.com/bwmarrin/discordgo"
)

func init() {
	var defaultMemberPermissions int64 = discordgo.PermissionViewChannel | discordgo.PermissionSendMessages
	dmPermission := false

	createCommand(&discordgo.ApplicationCommand{
		Name:        "categories",
		Description: "prints all categories",
		DefaultMemberPermissions: &defaultMemberPermissions,
		DMPermission:             &dmPermission,
	}, func(dg *discordgo.Session, i *discordgo.InteractionCreate) {
		// Create a string of all categories
		categories := strings.Join(news.Categories, ", ")

		// Create an embed message
		embed := &discordgo.MessageEmbed{
			Title:       "Categories",
			Description: categories,
			Color:       0x0099ff,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Daily News",
			},
		}

		// Send the embed as a response to the interaction
		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		}

		dg.InteractionRespond(i.Interaction, response)
	})
}
