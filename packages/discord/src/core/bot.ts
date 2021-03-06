import { Client, Intents, MessageEmbed } from "discord.js";
import { logger } from "@daily-news-discord/logger";
import { PREFIX } from "@daily-news-discord/environment";
import { handleMessage } from "./commandHandler";
import { sendNews } from "../utils/sendNews";

export const client = new Client({
  intents: [
    Intents.FLAGS.GUILD_MEMBERS,
    Intents.FLAGS.GUILDS,
    Intents.FLAGS.GUILD_MESSAGES,
    Intents.FLAGS.GUILD_MESSAGE_REACTIONS,
    Intents.FLAGS.GUILD_PRESENCES,
    Intents.FLAGS.GUILD_EMOJIS_AND_STICKERS,
  ],
});

client.on("error", (e) => {
  logger.error(e.message);
});

client.on("warn", (w) => {
  logger.warn(w);
});

client.on("debug", (d) => {
  logger.debug(d);
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

  sendNews();
});

client.on("messageCreate", async (message) => {
  handleMessage(message);
});

client.on("guildCreate", async (guild) => {
  const owner = await guild.fetchOwner();

  const embed = new MessageEmbed()
    .setTitle("Thank you for using News Bot!")
    .setDescription(
      `Hello ${owner.displayName}, thank you for adding me to your server!\n\n` +
        `To get started, use the command \`${PREFIX}help\` to see a list of commands.\n\n`,
    )
    .setFooter({ text: "Daily News" });

  await owner.user.send({ embeds: [embed] });
});
