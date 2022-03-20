import { NEWS_API_ENDPOINT } from "@daily-news-discord/environment";
import axios from "axios";

interface NewsAPIResponse {
  category: string;
  data: Datum[];
  success: boolean;
}

interface Datum {
  author: string;
  content: string;
  date: string;
  imageUrl: string;
  readMoreUrl: string;
  time: string;
  title: string;
  url: string;
}

const categories = [
  "all",
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

type Category = typeof categories[number];

export const getNews = async (category: Category) => {
  const response = await axios.get(`${NEWS_API_ENDPOINT}${category}`);
  const result: NewsAPIResponse = await response.data();

  return result;
};
