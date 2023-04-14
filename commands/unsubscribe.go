package commands

import (
	"log"
	"strings"

	"github.com/SushiWaUmai/daily-news-discord/db"
	"github.com/SushiWaUmai/daily-news-discord/news"
	"github.com/bwmarrin/discordgo"
)

func init() {
	var categoryChoices []*discordgo.ApplicationCommandOptionChoice
	for _, c := range news.Categories {
		categoryChoices = append(categoryChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  c,
			Value: c,
		})
	}
	categoryChoices = append(categoryChoices, &discordgo.ApplicationCommandOptionChoice{
		Name:  "all",
		Value: "all",
	})

	var defaultMemberPermissions int64 = discordgo.PermissionManageServer
	dmPermission := false

	createCommand(&discordgo.ApplicationCommand{
		Name:        "unsubscribe",
		Description: "Unsubscribe News to this channel",
		DefaultMemberPermissions: &defaultMemberPermissions,
		DMPermission:             &dmPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "category",
				Description: "The news category you want to unsubscribe from",
				Required:    true,
				Choices:     categoryChoices,
			},
		},
	}, func(dg *discordgo.Session, i *discordgo.InteractionCreate) {
		guildID := i.GuildID
		channelID := i.ChannelID
		category := strings.ToLower(i.ApplicationCommandData().Options[0].StringValue())

		if category == "all" {
			db.DeleteGuildChannels(guildID, channelID)

			dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You are now unsubscribed from all categories",
				},
			})
			return
		}

		guildChannel, _ := db.GetGuildChannel(guildID, channelID, category)
		if guildChannel == nil {
			dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You are already unsubscribed to this channel",
				},
			})
			return
		}

		err := db.DeleteGuildChannel(guildID, channelID, category)
		if err != nil {
			log.Println("Failed to delete GuildChannel", err)
			return
		}

		dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You are now unsubscribed to this channel",
			},
		})
	})
}
