import { prismaClient } from "../core";

export const createGuildChannel = async (
  guildId: string,
  channelId: string,
) => {
  await prismaClient.guildChannel.create({
    data: {
      guild_id: guildId,
      channel_id: channelId,
    },
  });
};

export const deleteGuildChannel = async (
  guildId: string,
  channelId: string,
) => {
  await prismaClient.guildChannel.deleteMany({
    where: {
      guild_id: guildId,
      channel_id: channelId,
    },
  });
};

export const getGuildChannel = async (guildId: string, channelId: string) => {
  return await prismaClient.guildChannel.findFirst({
    where: {
      guild_id: guildId,
      channel_id: channelId,
    },
  });
};
