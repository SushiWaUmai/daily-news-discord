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

export const deleteGuildChannel = async (
  guildId: string,
  channelId: string,
) => {
  await prismaClient.guildChannel.delete({
    where: {
      guild: guildId,
      channel: channelId,
    },
  });
};

export const getGuildChannel = async (guildId: string, channelId: string) => {
  return await prismaClient.guildChannel.findOne({
    where: {
      guild: guildId,
      channel: channelId,
    },
  });
};
