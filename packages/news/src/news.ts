import { NEWS_API_ENDPOINT } from "@daily-news-discord/environment";
import { logger } from "@daily-news-discord/logger";
import axios from "axios";

export interface NewsAPIResponse {
  category: string;
  data: Datum[];
  success: boolean;
}

export interface Datum {
  author: string;
  content: string;
  date: string;
  imageUrl: string;
  readMoreUrl: string;
  time: string;
  title: string;
  url: string;
}

export const categories = [
  "business",
  "sports",
  "world",
  "politics",
  "technology",
  "startup",
  "entertainment",
  "miscellaneous",
  "hatke",
  "science",
  "automobile",
] as const;

export const categoryArray = categories.map((category) => category as string);

export type Category = typeof categories[number];

export const getNews = async (category: Category) => {
  try {
    const response = await axios.get(`${NEWS_API_ENDPOINT}${category}`);
    const result: NewsAPIResponse = await response.data;
    return result;
  } catch (e) {
    logger.error("Error fetching news: ", e);
    return;
  }
};
