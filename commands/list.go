package commands

import (
	"strings"

	"github.com/SushiWaUmai/daily-news-discord/db"
	"github.com/bwmarrin/discordgo"
)

func init() {
	var defaultMemberPermissions int64 = discordgo.PermissionViewChannel | discordgo.PermissionSendMessages
	dmPermission := false

	createCommand(&discordgo.ApplicationCommand{
		Name:        "list",
		Description: "sends a list of subscribed news",
		DefaultMemberPermissions: &defaultMemberPermissions,
		DMPermission:             &dmPermission,
	}, func(dg *discordgo.Session, i *discordgo.InteractionCreate) {
		guildID := i.GuildID
		channelID := i.ChannelID

		categories, err := db.GetGuildChannels(guildID, channelID)

		if err != nil {
			dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "An error occurred while retrieving the subscribed categories.",
				},
			})
			return
		}

		if len(categories) == 0 {
			dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You are not subscribed to any categories.",
				},
			})
			return
		}

		var categoryNames []string
		for _, category := range categories {
			categoryNames = append(categoryNames, "`"+category.Category+"`")
		}

		embed := &discordgo.MessageEmbed{
			Title:       "Subscribed Categories",
			Description: strings.Join(categoryNames, ", "),
			Color:       0x0099ff,
		}

		dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	})
}
