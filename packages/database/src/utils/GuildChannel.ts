import { prismaClient } from "../core";

export const createGuildChannel = async (
  guildId: string,
  channelId: string,
  category: string,
) => {
  await prismaClient.guildChannel.create({
    data: {
      guild_id: guildId,
      channel_id: channelId,
      category: category,
    },
  });
};

export const createGuildChannels = async (
  guildId: string,
  channelId: string,
  categories: string[],
) => {
  // Convert to array of objects
  const data = categories.map((category) => ({
    guild_id: guildId,
    channel_id: channelId,
    category: category,
  }));

  await prismaClient.guildChannel.createMany({
    data: data,
    skipDuplicates: true,
  });
};

export const deleteGuildChannel = async (
  guildId: string,
  channelId: string,
  category: string,
) => {
  await prismaClient.guildChannel.deleteMany({
    where: {
      guild_id: guildId,
      channel_id: channelId,
      category: category,
    },
  });
};

export const deleteGuildChannels = async (
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

export const getGuildChannel = async (
  guildId: string,
  channelId: string,
  category?: string,
) => {
  return await prismaClient.guildChannel.findFirst({
    where: {
      guild_id: guildId,
      channel_id: channelId,
      category: category,
    },
  });
};

export const getGuildChannels = async (guildId: string, channelId: string) => {
  return await prismaClient.guildChannel.findMany({
    where: {
      guild_id: guildId,
      channel_id: channelId,
    },
  });
};
