import { Command } from "../utils/Command";
import { PREFIX } from "@daily-news-discord/environment";
import { commands } from "../core/commandHandler";
import { MessageEmbed } from "discord.js";

export default new Command({
  name: "help",
  description: "Shows a list of commands",
  aliases: ["commands", "cmds"],
  usage: `${PREFIX}help [command]`,
  execute: async (msg, args) => {
    // Create a Embed of all commands
    const embed = new MessageEmbed()
      .setTitle("Commands")
      .setColor("#0099ff")
      .setFooter({ text: "Daily News" });

    // If no command is specified, show all commands
    if (!args[0]) {
      embed.setDescription(commands.map((c) => `\`${c.name}\``).join(", "));
    }

    // If a command is specified, show that command
    else {
      const commandName = args[0].toLocaleLowerCase();
      const command = commands.find(
        (command) =>
          command.name === commandName ||
          command.aliases?.includes(commandName),
      );

      if (!command) {
        embed.setDescription(`\`${commandName}\` is not a valid command`);
      } else {
        embed.setDescription(`\`${command.name}\` - ${command.description}`);
      }
    }

    // Send the embed
    msg.reply({ embeds: [embed] });
  },
});
