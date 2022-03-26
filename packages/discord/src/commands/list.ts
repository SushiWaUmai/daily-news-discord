import { getGuildChannels } from "@daily-news-discord/database";
import { PREFIX } from "@daily-news-discord/environment";
import { MessageEmbed } from "discord.js";
import { Command } from "../utils/Command";

export default new Command({
  name: "list",
  description: "Lists all categories subscribed to the channel.",
  usage: `${PREFIX}list`,
  aliases: ["list-categories", "list-cats", "list-subs"],
  execute: async (message) => {
    const guild = message.guild;
    const channel = message.channel;

    if (!guild) {
      message.reply("This command can only be used in a server.");
      return;
    }

    const subscribed = await getGuildChannels(guild.id, channel.id);

    if (subscribed.length === 0) {
      message.reply("You are not subscribed to any categories.");
      return;
    }

    const embed = new MessageEmbed()
      .setTitle("Subscribed Categories")
      .setDescription(subscribed.map((c) => `\`${c.category}\``).join(", "))
      .setColor("#0099ff");

    await message.reply({ embeds: [embed] });
  },
});
