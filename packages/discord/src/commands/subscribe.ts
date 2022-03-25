import { PREFIX } from "@daily-news-discord/environment";
import { Command } from "../utils/Command";
import { categoryArray } from "@daily-news-discord/news";
import {
  createGuildChannel,
  getGuildChannel,
} from "@daily-news-discord/database";

export default new Command({
  name: "subscribe",
  description: "Subscribe News to this channel",
  aliases: ["sub"],
  usage: `${PREFIX}subscribe <category>`,
  execute: async (msg, args) => {
    const category = args[0];
    if (!category) {
      msg.reply("Please specify a category");
      return;
    }

    // Check if category is valid
    if (!categoryArray.includes(category)) {
      msg.reply("Please specify a valid category");
      return;
    }

    if (!msg.guild) {
      msg.reply("Please use this command in a server");
      return;
    }

    const channel = msg.channel;
    const guild = msg.guild;

    const guildChannel = await getGuildChannel(guild.id, channel.id);
    if (guildChannel) {
      msg.reply("You are already subscribed to this channel");
      return;
    }
    createGuildChannel(guild.id, channel.id);
    msg.reply(`You are now subscribed to ${category}`);
  },
});
