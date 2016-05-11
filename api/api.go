package api

const Endpoint = "https://api.findmespot.com/spot-main-web/consumer/rest-api/2.0/public/feed/"

type Spot struct {
	Response struct {
		FeedResponse struct {
			Count int `json:"count"`
			Feeds struct {
				Feed struct {
					ID                   string `json:"id"`
					Name                 string `json:"name"`
					Description          string `json:"description"`
					Status               string `json:"status"`
					Usage                int    `json:"usage"`
					DaysRange            int    `json:"daysRange"`
					DetailedMessageShown bool   `json:"detailedMessageShown"`
				} `json:"feed"`
			} `json:"feeds"`
		} `json:"feedResponse"`
	} `json:"response"`
}
