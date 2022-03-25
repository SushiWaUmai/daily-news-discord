import { Command } from "../utils/Command";
import { PREFIX } from "@daily-news-discord/environment";
import { MessageEmbed } from "discord.js";
import { categoryArray } from "@daily-news-discord/news";

export default new Command({
  name: "categories",
  description: "Shows a list of categories",
  aliases: ["cat"],
  usage: `${PREFIX}help [command]`,
  execute: async (msg) => {
    // Create a Embed of all categories
    const embed = new MessageEmbed()
      .setTitle("Categories")
      .setColor("#0099ff")
      .setFooter({ text: "Daily News" });

    embed.setDescription(categoryArray.map((c) => `\`${c}\``).join(", "));
    // Send the embed
    msg.reply({ embeds: [embed] });
  },
});
