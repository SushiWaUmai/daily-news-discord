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

	createCommand(&discordgo.ApplicationCommand{
		Name:        "subscribe",
		Description: "subscribes news to this channel",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "category",
				Description: "The news category you want to subscribe to",
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
			db.CreateGuildChannels(guildID, channelID, news.Categories)

			dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You are now subscribed to all categories",
				},
			})
			return
		}

		guildChannel, _ := db.GetGuildChannel(guildID, channelID, category)
		if guildChannel != nil {
			dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You are already subscribed to this category",
				},
			})
			return
		}

		err := db.CreateGuildChannel(guildID, channelID, category)
		if err != nil {
			log.Println("Failed to create GuildChannel", err)
			return
		}

		dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You are now subscribed to " + category,
			},
		})
	})
}
