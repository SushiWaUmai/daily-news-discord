package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	createCommand(&discordgo.ApplicationCommand{
		Name:        "help",
		Description: "shows a list of commands",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "command",
				Description: "The specified command where help is needed",
				Required:    false,
			}},
	}, func(dg *discordgo.Session, i *discordgo.InteractionCreate) {
		// Create an empty embed message
		embed := &discordgo.MessageEmbed{
			Title: "Commands",
			Color: 0x0099ff,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Daily News",
			},
		}

		// If no command is specified, show all commands
		if len(i.ApplicationCommandData().Options) == 0 {
			commandFields := make([]*discordgo.MessageEmbedField, len(CommandMap))
			for _, cmd := range CommandMap {
				commandFields = append(commandFields, &discordgo.MessageEmbedField{
					Name:   cmd.AppCmd.Name,
					Value:  cmd.AppCmd.Description,
				})
			}
			embed.Fields = commandFields
		} else {
			// If a command is specified, show that command
			commandName := strings.ToLower(i.ApplicationCommandData().Options[0].StringValue())
			command, ok := CommandMap[commandName]

			if ok {
				embed.Title = "Command " + command.AppCmd.Name
				embed.Description = "`" + command.AppCmd.Name + "` - " + command.AppCmd.Description
			} else {
				embed.Title = "Command not found"
				embed.Description = "`" + commandName + "` is not a valid command"
			}
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
