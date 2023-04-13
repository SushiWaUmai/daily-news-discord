package db

import (
	"fmt"
	"log"

	"github.com/SushiWaUmai/daily-news-discord/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GuildChannel struct {
	gorm.Model
	GuildID   string `gorm:"not null"`
	ChannelID string `gorm:"not null"`
	Category  string `gorm:"not null"`
}

var db *gorm.DB

func init() {
	var err error

	// Create postgres connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", env.DB_HOST, env.DB_USERNAME, env.DB_PASSWORD, env.DB_NAME, env.DB_PORT)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicln("Failed to open database connection:", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&GuildChannel{})
	if err != nil {
		log.Panicln("Failed to migrate database:", err)
	}
}

func CreateGuildChannel(guildID string, channelID string, category string) error {
	guildChannel := &GuildChannel{
		GuildID:   guildID,
		ChannelID: channelID,
		Category:  category,
	}
	return db.Create(guildChannel).Error
}

func CreateGuildChannels(guildID string, channelID string, categories []string) error {
	var guildChannels []GuildChannel
	for _, category := range categories {
		guildChannels = append(guildChannels, GuildChannel{
			GuildID:   guildID,
			ChannelID: channelID,
			Category:  category,
		})
	}
	return db.Create(&guildChannels).Error
}

func DeleteGuildChannel(guildID string, channelID string, category string) error {
	return db.Where("guild_id = ? AND channel_id = ? AND category = ?", guildID, channelID, category).Delete(&GuildChannel{}).Error
}

func DeleteGuildChannels(guildID string, channelID string) error {
	return db.Where("guild_id = ? AND channel_id = ?", guildID, channelID).Delete(&GuildChannel{}).Error
}

func GetGuildChannel(guildID string, channelID string, category string) (*GuildChannel, error) {
	var guildChannel GuildChannel
	err := db.Where("guild_id = ? AND channel_id = ? AND category = ?", guildID, channelID, category).First(&guildChannel).Error
	if err != nil {
		return nil, err
	}
	return &guildChannel, nil
}

func GetGuildChannels(guildID string, channelID string) ([]GuildChannel, error) {
	var guildChannels []GuildChannel
	err := db.Where("guild_id = ? AND channel_id = ?", guildID, channelID).Find(&guildChannels).Error
	if err != nil {
		return nil, err
	}
	return guildChannels, nil
}

func GetAllGuildChannels(db *gorm.DB) ([]GuildChannel, error) {
	var guildChannels []GuildChannel
	err := db.Find(&guildChannels).Error
	if err != nil {
		return nil, err
	}
	return guildChannels, nil
}
