package bot

import (
	"log"
	"math/rand"

	"github.com/SushiWaUmai/daily-news-discord/commands"
	"github.com/SushiWaUmai/daily-news-discord/db"
	"github.com/SushiWaUmai/daily-news-discord/news"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

type NewsBot struct {
	session *discordgo.Session
	cronJob *cron.Cron
}

func CreateBot(token string) (*NewsBot, error) {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}

	news := &NewsBot{
		session: dg,
		cronJob: cron.New(),
	}

	return news, err
}

func (bot *NewsBot) Start() error {
	bot.session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent
	bot.session.AddHandler(handleReady)
	bot.session.AddHandler(handleInteraction)

	err := bot.session.Open()
	if err != nil {
		return err
	}

	err = bot.StartJob()
	if err != nil {
		return err
	}

	return nil
}

func (bot *NewsBot) Close() {
	bot.cronJob.Stop()
	commands.UnregisterCommands(bot.session)
	bot.session.Close()
}

func (bot *NewsBot) StartJob() error {
	_, err := bot.cronJob.AddFunc("0 0 * * *", bot.sendNews)

	if err != nil {
		return err
	}

	log.Println("Starting Cron Job...")
	bot.cronJob.Start()

	return nil
}

func (bot *NewsBot) sendNews() {
	log.Println("Fetching news...")

	// get the news for all categories
	newsData := make(map[string]*news.NewsAPIResponse)
	for _, category := range news.Categories {
		n, err := news.GetNews(category)
		if err != nil {
			log.Printf("Error fetching news for category %s: %v\n", category, err)
			continue
		}
		log.Printf("Fetched news for category %s\n", category)
		newsData[category] = n
	}

	log.Println("Creating Embeds...")

	newsEmbeds := make(map[string][]*discordgo.MessageEmbed)
	for category, data := range newsData {
		var embeds []*discordgo.MessageEmbed
		for _, d := range data.Data {
			embed := &discordgo.MessageEmbed{
				Title:       d.Title,
				Description: d.Content,
				Color:       0x0099ff,
				URL:         d.ReadMoreURL,
				Image: &discordgo.MessageEmbedImage{
					URL: d.ImageURL,
				},
				Author: &discordgo.MessageEmbedAuthor{
					Name: d.Author,
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: category,
				},
			}
			embeds = append(embeds, embed)
		}
		newsEmbeds[category] = embeds
	}

	subbed, err := db.GetAllGuildChannels()

	if err != nil {
		log.Printf("Error getting subscribed channels: %v\n", err)
		return
	}

	log.Printf("Sending news to %d channels...\n", len(subbed))

	// Optimize Discord API calls
	for _, guildChannel := range subbed {
		channel, err := bot.session.Channel(guildChannel.ChannelID)
		if err != nil {
			log.Printf("Error finding channel %s for guild %s: %v\n", guildChannel.ChannelID, guildChannel.GuildID, err)
			continue
		}

		embeds := newsEmbeds[guildChannel.Category]
		if len(embeds) == 0 {
			continue
		}

		embed := embeds[rand.Intn(len(embeds))]

		_, err = bot.session.ChannelMessageSendEmbed(channel.ID, embed)
		if err != nil {
			log.Printf("Error sending message to channel %s for guild %s: %v\n", guildChannel.ChannelID, guildChannel.GuildID, err)
		}
	}
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
