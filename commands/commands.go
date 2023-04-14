package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var CommandMap map[string]Command = make(map[string]Command)

type Command struct {
	AppCmd  *discordgo.ApplicationCommand
	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func createCommand(cmd *discordgo.ApplicationCommand, execute func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	log.Println("Creating command", cmd.Name)
	CommandMap[cmd.Name] = Command{
		AppCmd:  cmd,
		Execute: execute,
	}
}

func RegisterCommands(dg *discordgo.Session) {
	for _, c := range CommandMap {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", c.AppCmd)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", c.AppCmd.Name, err)
		}
	}
}

func UnregisterCommands(dg *discordgo.Session) {
	for _, c := range CommandMap {
		err := dg.ApplicationCommandDelete(dg.State.User.ID, "", c.AppCmd.ID)
		if err != nil {
			log.Printf("Could not delete '%v' command: %v", c.AppCmd.Name, err)
		}
	}
}
