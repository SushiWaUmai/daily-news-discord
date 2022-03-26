import { PREFIX } from "@daily-news-discord/environment";
import { Command } from "../utils/Command";
import {
  deleteGuildChannel,
  deleteGuildChannels,
  getGuildChannel,
  getGuildChannels,
} from "@daily-news-discord/database";
import { categoryArray } from "@daily-news-discord/news";

export default new Command({
  name: "unsubscribe",
  description: "Unsubscribe News to this channel",
  aliases: ["unsub"],
  usage: `${PREFIX}unsubscribe <category>`,
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
      const guildChannels = await getGuildChannels(guild.id, channel.id);
      if (guildChannels.length === 0) {
        msg.reply("You are not subscribed to any categories");
        return;
      }

      await deleteGuildChannels(guild.id, channel.id);
      msg.reply("You are now unsubscribed from all categories");
      return;
    }

    const guildChannel = await getGuildChannel(guild.id, channel.id, category);
    if (!guildChannel) {
      msg.reply("You are already unsubscribed to this channel");
      return;
    }

    await deleteGuildChannel(guild.id, channel.id, category);
    msg.reply("You are now unsubscribed to this channel");
  },
});
