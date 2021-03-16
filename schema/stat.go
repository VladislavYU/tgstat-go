package schema

type StatResponse struct {
	Status   string `json:"status"`
	Response []struct {
		ServiceKey    string `json:"serviceKey"`
		Title         string `json:"title"`
		SpentChannels string `json:"spentChannels,omitempty"`
		SpentRequests string `json:"spentRequests"`
		ExpiredAt     int    `json:"expiredAt"`
		SpentWords    string `json:"spentWords,omitempty"`
	} `json:"response"`
}