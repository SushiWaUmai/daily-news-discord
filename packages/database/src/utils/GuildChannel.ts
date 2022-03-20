import { prismaClient } from "../core";

export const createGuildChannel = async (
  guildId: string,
  channelId: string,
) => {
  await prismaClient.guildChannel.create({
    data: {
      guild: guildId,
      channel: channelId,
    },
  });
};
