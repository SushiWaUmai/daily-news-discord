package news

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	newsAPIEndpoint = "https://inshortsapi.vercel.app/news?category="
)

var Categories = []string{
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
}

type Datum struct {
	Author      string `json:"author"`
	Content     string `json:"content"`
	Date        string `json:"date"`
	ImageURL    string `json:"imageUrl"`
	ReadMoreURL string `json:"readMoreUrl"`
	Time        string `json:"time"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

type NewsAPIResponse struct {
	Category string  `json:"category"`
	Data     []Datum `json:"data"`
	Success  bool    `json:"success"`
}

func getNews(category string) (*NewsAPIResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", newsAPIEndpoint, category))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result NewsAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
