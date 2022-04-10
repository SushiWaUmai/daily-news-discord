// import { getGuildChannel } from "@daily-news-discord/database";
import { getAllGuildChannels } from "@daily-news-discord/database";
import { logger } from "@daily-news-discord/logger";
import { getNews, NewsAPIResponse } from "@daily-news-discord/news";
import { categories } from "@daily-news-discord/news";
import { MessageEmbed, TextChannel } from "discord.js";
import { client } from "../core/bot";

export const nextTime = () => {
  const now = new Date();
  // Every day
  const next = new Date(now.getFullYear(), now.getMonth(), now.getDate() + 1);

  return next;
};

export const sendNews = async () => {
  // wait until next time
  const next = nextTime();
  const now = new Date();
  const diff = next.getTime() - now.getTime();
  const wait = Math.max(diff, 0);

  logger.info(`Sleeping for ${wait}ms`);
  await new Promise((resolve) => setTimeout(resolve, wait));

  logger.info("Fetching news...");

  // get the news for all categories
  const newsData: { [id: string]: NewsAPIResponse } = {};
  for (const category of categories) {
    const news = await getNews(category);
    if (!news) continue;
    logger.info(`Fetched news for category ${category}`);
    newsData[category] = news;
  }

  logger.info("Creating Embeds...");

  const newsEmbeds: { [id: string]: MessageEmbed[] } = Object.entries(
    newsData,
  ).reduce((a, [category, data]) => {
    const embeds: MessageEmbed[] = data.data.map((d) => {
      return new MessageEmbed()
        .setTitle(d.title)
        .setDescription(d.content)
        .setColor("#0099ff")
        .setURL(d.readMoreUrl)
        .setImage(d.imageUrl)
        .setAuthor({ name: d.author });
    });

    return { ...a, [category]: embeds };
  }, {});

  const subbed = await getAllGuildChannels();
  logger.info(`Sending news to ${subbed.length} channels...`);

  // Optimize Discord API calls
  for (const guildChannel of subbed) {
    const channel = client.channels.cache.get(
      guildChannel.channel_id,
    ) as TextChannel;

    if (!channel) {
      logger.error(
        `Could not find channel ${guildChannel.channel_id} for guild ${guildChannel.guild_id}`,
      );
      continue;
    }

    const embeds = newsEmbeds[guildChannel.category];

    if (!embeds || embeds.length <= 0) {
      continue;
    }

    const embed = embeds[Math.floor(Math.random() * embeds.length)];

    await channel.send({ embeds: [embed] });
  }

  sendNews();
};
