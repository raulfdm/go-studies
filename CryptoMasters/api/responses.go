package api

type CEXResponse struct {
	Ok   string              `json:"ok"`
	Data map[string]Currency `json:"data"`
}

type Currency struct {
	BestBid                 string `json:"bestBid"`
	BestAsk                 string `json:"bestAsk"`
	BestBidChange           string `json:"bestBidChange"`
	BestBidChangePercentage string `json:"bestBidChangePercentage"`
	BestAskChange           string `json:"bestAskChange"`
	BestAskChangePercentage string `json:"bestAskChangePercentage"`
	Low                     string `json:"low"`
	High                    string `json:"high"`
	Volume30d               string `json:"volume30d"`
	LastTradeDateISO        string `json:"lastTradeDateISO"`
	Volume                  string `json:"volume"`
	QuoteVolume             string `json:"quoteVolume"`
	LastTradeVolume         string `json:"lastTradeVolume"`
	VolumeUSD               string `json:"volumeUSD"`
	Last                    string `json:"last"`
	LastTradePrice          string `json:"lastTradePrice"`
	PriceChange             string `json:"priceChange"`
	PriceChangePercentage   string `json:"priceChangePercentage"`
}
