import { client } from "@daily-news-discord/discord";
import { BOT_TOKEN } from "@daily-news-discord/environment";

const main = async () => {
  await client.login(BOT_TOKEN);
};

main();
