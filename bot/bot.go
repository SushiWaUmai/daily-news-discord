package bot

import (
	"log"

	"github.com/SushiWaUmai/daily-news-discord/commands"
	"github.com/bwmarrin/discordgo"
)

type NewsBot struct {
	session *discordgo.Session
}

func CreateBot(token string) (*NewsBot, error) {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}

	news := &NewsBot{
		session: dg,
	}

	return news, err
}

func (bot *NewsBot) Start() error {
	bot.session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent
	bot.session.AddHandler(handleReady)
	bot.session.AddHandler(handleInteraction)

	err := bot.session.Open()
	return err
}

func (bot *NewsBot) Close() {
	// commands.UnregisterCommands(bot.session)
	bot.Close()
}

func handleReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	commands.RegisterCommands(s)
}

func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if c, ok := commands.CommandMap[i.ApplicationCommandData().Name]; ok {
		c.Execute(s, i)
	}
}
