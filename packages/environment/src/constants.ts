import dotenv from "dotenv";

dotenv.config();

export const BuildMode = (process.env.NODE_ENV || "development") as
  | "development"
  | "production";

export const __prod__ = BuildMode === "production";
export const NEWS_API_ENDPOINT =
  "https://inshortsapi.vercel.app/news?category=";
