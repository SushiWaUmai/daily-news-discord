import dotenv from "dotenv";

dotenv.config({ path: `${__dirname}/../../../.env` });

export const BuildMode = (process.env.NODE_ENV || "development") as
  | "development"
  | "production";

export const __prod__ = BuildMode === "production";
export const BOT_TOKEN = process.env.BOT_TOKEN;

export const NEWS_API_ENDPOINT =
  "https://inshortsapi.vercel.app/news?category=";

export const PREFIX = "news ";
