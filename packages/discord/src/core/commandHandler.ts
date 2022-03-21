import { logger } from "@daily-news-discord/logger";
import { Command } from "../utils/Command";
import { Message, TextChannel } from "discord.js";
import fs from "fs";
import { PREFIX } from "@daily-news-discord/environment";

export const commands: Command[] = [];

async function addCommandsRecursive(dir: string, folder: string) {
  //recursion to scan directories inside "commands" too for better structure
  const commandFiles = fs
    .readdirSync(dir)
    .filter(
      (file) =>
        file.endsWith(".js") || fs.lstatSync(dir + "/" + file).isDirectory(),
    ); // only files that end with .js or folders
  for (const file of commandFiles) {
    if (fs.lstatSync(dir + "/" + file.toString()).isDirectory()) {
      logger.info(`Registering category ${file}`);
      await addCommandsRecursive(dir + "/" + file.toString(), file);
    }
    if (file.endsWith(".js")) {
      const command = (await import(`../commands/${folder}/${file}`))
        .default as Command | undefined;
      if (command?.name) {
        commands.push(command);
        logger.info(`Registered ${command.name} command ${dir}/${file}`);
      } else {
        logger.warn(`Command in ${dir}/${file} is not a valid command`);
      }
    }
  }
}

addCommandsRecursive(`${__dirname}/../commands`, "");

export const handleMessage = async (message: Message) => {
  if (!message.guild || message.webhookId || message.author.bot) {
    return;
  }

  let content = message.content;

  if (content.toLocaleLowerCase().startsWith(PREFIX)) {
    content = content.slice(PREFIX.length);
    const args = content.split(" ");
    const commandName = args[0].toLocaleLowerCase();
    args.shift();

    for (const command of commands) {
      if (
        command.name === commandName ||
        (command.aliases && command.aliases.includes(commandName))
      ) {
        logger.info(`Executing Command ${command.name} with args [${args}]`);

        try {
          await command.execute(message, args);
        } catch (e) {
          logger.error(e);
          if (
            message.guild &&
            message.guild.me &&
            message.channel.type == "GUILD_TEXT"
          ) {
            if (
              message.guild.me
                .permissionsIn(message.channel as TextChannel)
                .has("SEND_MESSAGES")
            )
              message.reply(
                `An error occured while executing that command. Please contact the developer. Or try again later. Error: ${e.message}`,
              );
          }
        }

        return;
      }
    }
  }
};
