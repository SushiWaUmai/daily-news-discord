import { PREFIX } from "@daily-news-discord/environment";
import { Command } from "../utils/Command";
import { categoryArray } from "@daily-news-discord/news";
import {
  createGuildChannel,
  createGuildChannels,
  deleteGuildChannels,
  getGuildChannel,
} from "@daily-news-discord/database";

export default new Command({
  name: "subscribe",
  description: "Subscribe News to this channel",
  aliases: ["sub"],
  usage: `${PREFIX}subscribe <category>`,
  execute: async (msg, args) => {
    if (!msg.guild) {
      msg.reply("Please use this command in a server");
      return;
    }

    const category = args[0];
    if (!category) {
      msg.reply("Please specify a category");
      return;
    }

    // Check if category is valid
    if (!categoryArray.includes(category) && category !== "all") {
      msg.reply("Please specify a valid category");
      return;
    }

    const channel = msg.channel;
    const guild = msg.guild;

    if (category === "all") {
      await deleteGuildChannels(guild.id, channel.id);
      await createGuildChannels(guild.id, channel.id, categoryArray);
      msg.reply("You are now subscribed to all categories");
      return;
    }

    const guildChannel = await getGuildChannel(guild.id, channel.id, category);
    if (guildChannel) {
      msg.reply("You are already subscribed to this channel");
      return;
    }

    await createGuildChannel(guild.id, channel.id, category);
    msg.reply(`You are now subscribed to ${category}`);
  },
});
