import { Client, MessageEmbed } from "discord.js";
import { logger } from "@daily-news-discord/logger";
import { PREFIX } from "@daily-news-discord/environment";

export const client = new Client({
  intents: [],
});

client.on("error", (e) => {
  logger.error(e.message);
});
client.on("warn", (w) => {
  logger.warn(w);
});

client.on("ready", async () => {
  logger.info(`Logged in as ${client.user?.username}`);

  client.user?.setPresence({
    status: "online",
    activities: [
      {
        name: "donda",
        type: "LISTENING",
      },
    ],
  });
});

client.on("messageCreate", async (message) => {
  logger.info(message.content);
});

client.on("guildCreate", async (guild) => {
  const owner = await guild.fetchOwner();

  const embed = new MessageEmbed()
    .setTitle("Thank you for using GoBot!")
    .setDescription(
      `Hello ${owner.displayName}, thank you for adding me to your server!\n\n` +
        `To get started, use the command \`${PREFIX}help\` to see a list of commands.\n\n`,
    )
    .setFooter({ text: "Daily News" });

  await owner.user.send({ embeds: [embed] });
});
