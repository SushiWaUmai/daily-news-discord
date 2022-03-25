import { PREFIX } from "@daily-news-discord/environment";
import { Command } from "../utils/Command";
import {
  deleteGuildChannel,
  getGuildChannel,
} from "@daily-news-discord/database";

export default new Command({
  name: "unsubscribe",
  description: "Unsubscribe News to this channel",
  aliases: ["unsub"],
  usage: `${PREFIX}unsubscribe <category>`,
  execute: async (msg) => {
    if (!msg.guild) {
      msg.reply("Please use this command in a server");
      return;
    }

    const channel = msg.channel;
    const guild = msg.guild;

    const guildChannel = await getGuildChannel(guild.id, channel.id);
    if (!guildChannel) {
      msg.reply("You are already unsubscribed to this channel");
      return;
    }
    deleteGuildChannel(guild.id, channel.id);
    msg.reply("You are now unsubscribed to this channel");
  },
});
